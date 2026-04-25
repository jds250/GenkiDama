// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {RLPReader} from "solidity-rlp/contracts/RLPReader.sol";

/* Internal Imports */
import {DataTypes} from "./DataTypes.sol";
import {IZKVerifier} from "./IZKVerifier.sol";
import {HeaderLib} from "./HeaderLib.sol";
import {MPTVerifier} from "./MPTVerifier.sol";

interface IBCPRollupTarget {
    function isBlockInDispute(uint256 _blockNumber) external view returns (bool);
    function getBlockCommitmentHash(uint256 _blockNumber) external view returns (bytes32);
    function rollBackBlock(uint256 _targetBlockHeight) external;
}

/// @title 挑战状态管理
// 专门用于管理挑战状态
contract BCPManager {
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;

    /*** Fields ***/
    /* Contract Instance */
    // The Data Type library
    DataTypes myDT;
    // The Data Type library
    IZKVerifier myZKV;
    MPTVerifier myMPTV;

    /* Const */
    bytes32 public constant ZERO_BYTES32 = 0x0000000000000000000000000000000000000000000000000000000000000000;
    // TODO: Set a reasonable wait period
    uint256 constant WAIT_PERIOD = 4;
    uint256 constant CONFIRM_PERIOD = 10;
    uint256 constant ROLLUP_BLOCKS_SLOT = 6;
    uint256 constant ROLLUP_BLOCK_FIELD_COUNT = 6;
    uint256 constant ROLLUP_COMMITMENT_FIELD_COUNT = 5;

    /* rollup system state */
    // All the blocks!
    uint256 public chainID;
    mapping(string=>DataTypes.ChallengeRecord) private challengePool;

    // checkpoint record
    mapping(string=>DataTypes.ChainCheckPoint) private checkPointPool;
    mapping(uint256=>DataTypes.ChainCheckPoint) public chainCheckPoints;

    // authority
    address[] public committerAddresses;
    address public superAddress;

    /* Events */
    event ChallengeStateNotify(string index, uint256 currentState, string desc);
    event BlockInfoNotify(string index, uint256 style, string desc);

    /***************
     * Constructor *
     **************/
    constructor(
        uint256 _chainID,
        address _dataTypesAddress,
        address _zkVerifierAddress,
        address _mptVerifierAddress
    ) {
        chainID = _chainID;
        myDT = DataTypes(_dataTypesAddress);
        myZKV = IZKVerifier(_zkVerifierAddress);
        myMPTV = MPTVerifier(_mptVerifierAddress);
        superAddress = msg.sender;
    }

    function GetChallengeMeta(string memory _index) public view returns (
        DataTypes.ChallengeState memory detail,
        uint256 confirmTime,
        address challenger,
        address questioner,
        uint256 state,
        uint256 begin,
        uint256 end
    ) {
        DataTypes.ChallengeRecord storage target = challengePool[_index];
        return (
            target.detail,
            target.confirmTime,
            target.challenger,
            target.questioner,
            target.state,
            target.begin,
            target.end
        );
    }

    function GetChallengeBlocks(string memory _index) public view returns (
        DataTypes.L1BlockInfo memory beginBlock,
        DataTypes.L1BlockInfo memory endBlock,
        DataTypes.L1BlockInfo memory middleBlock
    ) {
        DataTypes.ChallengeRecord storage target = challengePool[_index];
        return (target.beginBlock, target.endBlock, target.middleBlock);
    }

    function GetCheckPoint(string memory _index) public view returns (DataTypes.ChainCheckPoint memory) {
        return checkPointPool[_index];
    }

    modifier onlySuper() {
        require(
            msg.sender == address(superAddress),
            "Only super admin may perform action"
        );
        _;
    }

    function _isTerminalState(uint256 _state) internal pure returns (bool) {
        return _state == 5 || _state == 6;
    }

    function _isChallengeLive(DataTypes.ChallengeRecord storage _record) internal view returns (bool) {
        return _record.state != 0 && !_isTerminalState(_record.state);
    }

    function _decodeBytes32(bytes memory _value) internal pure returns (bytes32 result) {
        require(_value.length == 32, "invalid bytes32");
        assembly {
            result := mload(add(_value, 32))
        }
    }

    function _decodeBytes32At(bytes memory _value, uint256 _offset) internal pure returns (bytes32 result) {
        require(_value.length >= _offset + 32, "invalid bytes32 offset");
        assembly {
            result := mload(add(add(_value, 32), _offset))
        }
    }

    function _decodeUint256(bytes memory _value) internal pure returns (uint256) {
        return abi.decode(_value, (uint256));
    }

    function _tryDecodeUint256(bytes memory _value) internal pure returns (bool, uint256) {
        if (_value.length != 32) {
            return (false, 0);
        }
        return (true, abi.decode(_value, (uint256)));
    }

    function _decodeStateRootBundle(
        bytes memory _value
    ) internal pure returns (bool hasHeaderRoot, bytes32 headerRoot, bool hasCommitmentRoot, bytes32 commitmentRoot) {
        if (_value.length >= 32) {
            hasHeaderRoot = true;
            headerRoot = _decodeBytes32At(_value, 0);
        }

        if (_value.length >= 64) {
            hasCommitmentRoot = true;
            commitmentRoot = _decodeBytes32At(_value, 32);
        }
    }

    function _verifyConsistencyProof(
        bytes32 _remoteCommitmentHash,
        bytes32 _commitmentRoot,
        bytes[] memory _stateProof
    ) internal pure returns (bool) {
        if (_stateProof.length < 2) {
            return false;
        }

        uint256 path = _decodeUint256(_stateProof[1]);
        bytes32 computed = keccak256(abi.encodePacked(_remoteCommitmentHash));
        for (uint256 i = 2; i < _stateProof.length; i++) {
            if (_stateProof[i].length != 32) {
                return false;
            }

            bytes32 sibling = _decodeBytes32(_stateProof[i]);
            if (((path >> (i - 2)) & uint256(1)) == 0) {
                computed = keccak256(abi.encodePacked(computed, sibling));
            } else {
                computed = keccak256(abi.encodePacked(sibling, computed));
            }
        }

        return computed == _commitmentRoot;
    }

    function _decodeRlpBytes(bytes memory _encoded) internal pure returns (bytes memory) {
        return _encoded.toRlpItem().toBytes();
    }

    function _decodeRlpBytes32(bytes memory _encoded) internal pure returns (bytes32 result) {
        if (_encoded.length == 0) {
            return ZERO_BYTES32;
        }

        bytes memory raw = _decodeRlpBytes(_encoded);
        require(raw.length <= 32, "rlp bytes32 too large");
        assembly {
            result := mload(add(raw, 32))
        }
        if (raw.length < 32) {
            result = result >> (8 * (32 - raw.length));
        }
    }

    function _decodeAccountStorageRoot(bytes memory _accountValue) internal pure returns (bytes32) {
        RLPReader.RLPItem[] memory accountFields = _accountValue.toRlpItem().toList();
        require(accountFields.length == 4, "invalid account rlp");
        return bytes32(accountFields[2].toUintStrict());
    }

    function _decodeRlpUint256(bytes memory _encoded) internal pure returns (uint256) {
        if (_encoded.length == 0) {
            return 0;
        }
        return uint256(_decodeRlpBytes32(_encoded));
    }

    function _decodeAccountProofData(
        bytes[] memory _stateProof
    ) internal pure returns (
        bool hasAccountProof,
        bytes memory accountValue,
        bytes memory storageValue,
        bytes memory storageKey,
        uint256 accountProofCount,
        uint256 storageProofCount,
        uint256 siblingsOffset
    ) {
        if (_stateProof.length < 7) {
            return (false, bytes(""), bytes(""), bytes(""), 0, 0, 0);
        }

        accountValue = _stateProof[2];
        storageValue = _stateProof[3];
        storageKey = _stateProof[4];
        (bool ok1, uint256 decodedAccountProofCount) = _tryDecodeUint256(_stateProof[5]);
        (bool ok2, uint256 decodedStorageProofCount) = _tryDecodeUint256(_stateProof[6]);
        if (!ok1 || !ok2) {
            return (false, bytes(""), bytes(""), bytes(""), 0, 0, 0);
        }
        accountProofCount = decodedAccountProofCount;
        storageProofCount = decodedStorageProofCount;
        siblingsOffset = 7 + accountProofCount + storageProofCount;
        if (_stateProof.length < siblingsOffset) {
            return (false, bytes(""), bytes(""), bytes(""), 0, 0, 0);
        }

        return (true, accountValue, storageValue, storageKey, accountProofCount, storageProofCount, siblingsOffset);
    }

    function _verifyAccountStorageProof(
        bytes32 _headerStateRoot,
        uint256 _remoteChainID,
        bytes[] memory _stateProof
    ) internal view returns (bool, bytes32) {
        (
            bool hasAccountProof,
            bytes memory accountValue,
            bytes memory storageValue,
            bytes memory storageKey,
            uint256 accountProofCount,
            uint256 storageProofCount,
            uint256 siblingsOffset
        ) = _decodeAccountProofData(_stateProof);

        if (!hasAccountProof) {
            return (false, ZERO_BYTES32);
        }

        bytes[] memory accountProof = new bytes[](accountProofCount);
        for (uint256 i = 0; i < accountProofCount; i++) {
            accountProof[i] = _stateProof[7 + i];
        }

        bytes[] memory storageProof = new bytes[](storageProofCount);
        for (uint256 i = 0; i < storageProofCount; i++) {
            storageProof[i] = _stateProof[7 + accountProofCount + i];
        }

        DataTypes.ChainCheckPoint memory remoteCheckPoint = chainCheckPoints[_remoteChainID];
        bytes32 accountKey = keccak256(abi.encodePacked(remoteCheckPoint.contractAddr));
        MPTVerifier.MerkleProof memory accountMerkleProof = MPTVerifier.MerkleProof({
            expectedRoot: _headerStateRoot,
            key: abi.encodePacked(accountKey),
            proof: accountProof,
            keyIndex: 0,
            proofIndex: 0,
            expectedValue: accountValue
        });

        if (!myMPTV.verifyProof(accountMerkleProof)) {
            return (false, ZERO_BYTES32);
        }

        bytes32 storageRoot = _decodeAccountStorageRoot(accountValue);
        MPTVerifier.MerkleProof memory storageMerkleProof = MPTVerifier.MerkleProof({
            expectedRoot: storageRoot,
            key: storageKey,
            proof: storageProof,
            keyIndex: 0,
            proofIndex: 0,
            expectedValue: storageValue
        });
        if (!myMPTV.verifyProof(storageMerkleProof)) {
            return (false, ZERO_BYTES32);
        }

        if (_stateProof.length == siblingsOffset) {
            return (true, _decodeRlpBytes32(storageValue));
        }

        return (true, _decodeRlpBytes32(storageValue));
    }

    function _decodeDirectBlockProofData(
        bytes[] memory _stateProof
    ) internal pure returns (
        bool hasDirectProof,
        uint256 remoteRollupBlockID,
        bytes memory accountValue,
        uint256 accountProofCount,
        uint256 fieldProofCount,
        uint256 proofsOffset
    ) {
        if (_stateProof.length < 5) {
            return (false, 0, bytes(""), 0, 0, 0);
        }

        (bool okBlockID, uint256 decodedBlockID) = _tryDecodeUint256(_stateProof[1]);
        (bool okAccountProofCount, uint256 decodedAccountProofCount) = _tryDecodeUint256(_stateProof[3]);
        (bool okFieldProofCount, uint256 decodedFieldProofCount) = _tryDecodeUint256(_stateProof[4]);
        if (!okBlockID || !okAccountProofCount || !okFieldProofCount) {
            return (false, 0, bytes(""), 0, 0, 0);
        }
        remoteRollupBlockID = decodedBlockID;
        accountValue = _stateProof[2];
        accountProofCount = decodedAccountProofCount;
        fieldProofCount = decodedFieldProofCount;
        proofsOffset = 5 + accountProofCount;
        if (_stateProof.length < proofsOffset || fieldProofCount == 0) {
            return (false, 0, bytes(""), 0, 0, 0);
        }

        return (true, remoteRollupBlockID, accountValue, accountProofCount, fieldProofCount, proofsOffset);
    }

    function _rollupBlockFieldSlot(uint256 _blockID, uint256 _fieldIndex) internal pure returns (uint256) {
        return uint256(keccak256(abi.encode(uint256(ROLLUP_BLOCKS_SLOT)))) + (_blockID * ROLLUP_BLOCK_FIELD_COUNT) + _fieldIndex;
    }

    function _verifyDirectRollupBlockProof(
        bytes32 _headerStateRoot,
        uint256 _remoteChainID,
        uint256 _expectedRollupBlockID,
        bytes[] memory _stateProof
    ) internal view returns (bool, bytes32) {
        (
            bool hasDirectProof,
            uint256 remoteRollupBlockID,
            bytes memory accountValue,
            uint256 accountProofCount,
            uint256 fieldProofCount,
            uint256 cursor
        ) = _decodeDirectBlockProofData(_stateProof);

        if (!hasDirectProof || fieldProofCount != ROLLUP_COMMITMENT_FIELD_COUNT) {
            return (false, ZERO_BYTES32);
        }
        if (remoteRollupBlockID != _expectedRollupBlockID) {
            return (false, ZERO_BYTES32);
        }

        bytes[] memory accountProof = new bytes[](accountProofCount);
        for (uint256 i = 0; i < accountProofCount; i++) {
            accountProof[i] = _stateProof[5 + i];
        }

        DataTypes.ChainCheckPoint memory remoteCheckPoint = chainCheckPoints[_remoteChainID];
        bytes32 accountKey = keccak256(abi.encodePacked(remoteCheckPoint.contractAddr));
        MPTVerifier.MerkleProof memory accountMerkleProof = MPTVerifier.MerkleProof({
            expectedRoot: _headerStateRoot,
            key: abi.encodePacked(accountKey),
            proof: accountProof,
            keyIndex: 0,
            proofIndex: 0,
            expectedValue: accountValue
        });
        if (!myMPTV.verifyProof(accountMerkleProof)) {
            return (false, ZERO_BYTES32);
        }

        bytes32 storageRoot = _decodeAccountStorageRoot(accountValue);
        bytes32[ROLLUP_COMMITMENT_FIELD_COUNT] memory values;
        for (uint256 fieldIndex = 0; fieldIndex < ROLLUP_COMMITMENT_FIELD_COUNT; fieldIndex++) {
            if (_stateProof.length < cursor + 3) {
                return (false, ZERO_BYTES32);
            }

            bytes memory storageKey = _stateProof[cursor];
            bytes memory storageValue = _stateProof[cursor + 1];
            uint256 storageProofCount = _decodeUint256(_stateProof[cursor + 2]);
            cursor += 3;
            if (_stateProof.length < cursor + storageProofCount) {
                return (false, ZERO_BYTES32);
            }

            bytes32 expectedSlotKey = keccak256(abi.encode(_rollupBlockFieldSlot(remoteRollupBlockID, fieldIndex)));
            if (storageKey.length != 32 || _decodeBytes32(storageKey) != expectedSlotKey) {
                return (false, ZERO_BYTES32);
            }

            bytes[] memory storageProof = new bytes[](storageProofCount);
            for (uint256 proofIndex = 0; proofIndex < storageProofCount; proofIndex++) {
                storageProof[proofIndex] = _stateProof[cursor + proofIndex];
            }
            cursor += storageProofCount;

            MPTVerifier.MerkleProof memory storageMerkleProof = MPTVerifier.MerkleProof({
                expectedRoot: storageRoot,
                key: storageKey,
                proof: storageProof,
                keyIndex: 0,
                proofIndex: 0,
                expectedValue: storageValue
            });
            if (!myMPTV.verifyProof(storageMerkleProof)) {
                return (false, ZERO_BYTES32);
            }

            values[fieldIndex] = _decodeRlpBytes32(storageValue);
        }

        bytes32 commitmentHash = keccak256(
            abi.encode(
                uint256(values[0]),
                values[1],
                values[2],
                values[3],
                uint256(values[4])
            )
        );
        return (true, commitmentHash);
    }

    function _verifyHeaderBinding(bytes memory _expectedHeaderHash, bytes memory _header) internal pure returns (bool) {
        if (_expectedHeaderHash.length != 32 || _header.length == 0) {
            return false;
        }

        return HeaderLib.canonicalHash(_header) == _decodeBytes32(_expectedHeaderHash);
    }

    function _headerHashToProofInputs(bytes32 _headerHash) internal pure returns (uint256[4] memory inputs) {
        uint256 value = uint256(_headerHash);
        uint256 mask = uint256(type(uint64).max);

        inputs[0] = value & mask;
        inputs[1] = (value >> 64) & mask;
        inputs[2] = (value >> 128) & mask;
        inputs[3] = (value >> 192) & mask;
    }

    function _getRemoteCommitmentHash(DataTypes.ChallengeState memory _challenge) internal pure returns (bytes32, bool) {
        if (_challenge.stateProof.length == 0 || _challenge.stateProof[0].length != 32) {
            return (ZERO_BYTES32, false);
        }

        return (_decodeBytes32(_challenge.stateProof[0]), true);
    }

    function _verifyChallengeRequest(
        DataTypes.ChallengeState memory _challenge,
        bytes memory _header
    ) internal view returns (bool, string memory) {
        if (_challenge.account == address(0)) {
            return (false, "missing rollup address");
        }

        if (!_verifyHeaderBinding(_challenge.l1BlockData, _header)) {
            return (false, "header mismatch");
        }

        (bool hasHeaderRoot, bytes32 headerRoot, bool hasCommitmentRoot, bytes32 commitmentRoot) = _decodeStateRootBundle(_challenge.l1StateRoot);

        if (hasHeaderRoot) {
            (bool hasStateRoot, bytes32 headerStateRoot) = HeaderLib.tryStateRoot(_header);
            if (hasStateRoot && headerStateRoot != headerRoot) {
                return (false, "header state root mismatch");
            }
        }

        IBCPRollupTarget targetRollup = IBCPRollupTarget(_challenge.account);
        if (!targetRollup.isBlockInDispute(_challenge.l2BlockID)) {
            return (false, "challenge window closed");
        }

        (bytes32 remoteCommitmentHash, bool ok) = _getRemoteCommitmentHash(_challenge);
        if (!ok) {
            return (false, "missing commitment proof");
        }

        (bool hasDirectProof,,,,,) = _decodeDirectBlockProofData(_challenge.stateProof);
        if (hasDirectProof && !hasCommitmentRoot) {
            if (!hasHeaderRoot) {
                return (false, "missing header state root");
            }

            (bool directProofValid, bytes32 derivedCommitmentHash) = _verifyDirectRollupBlockProof(
                headerRoot,
                _challenge.chainID,
                _challenge.l2BlockID,
                _challenge.stateProof
            );
            if (!directProofValid) {
                return (false, "invalid rollup storage proof");
            }
            if (derivedCommitmentHash != remoteCommitmentHash) {
                return (false, "remote commitment mismatch");
            }
        } else {
            uint256 siblingsOffset = 2;
            (bool hasAccountProof,,,,,,uint256 parsedSiblingsOffset) = _decodeAccountProofData(_challenge.stateProof);
            if (hasAccountProof) {
                siblingsOffset = parsedSiblingsOffset;
                if (!hasHeaderRoot) {
                    return (false, "missing header state root");
                }

                (bool mptProofValid, bytes32 derivedCommitmentRoot) = _verifyAccountStorageProof(headerRoot, _challenge.chainID, _challenge.stateProof);
                if (!mptProofValid) {
                    return (false, "invalid state mpt proof");
                }
                if (hasCommitmentRoot && commitmentRoot != derivedCommitmentRoot) {
                    return (false, "commitment root mismatch");
                }
                commitmentRoot = derivedCommitmentRoot;
                hasCommitmentRoot = true;
            }

            if (hasCommitmentRoot) {
                bytes[] memory commitmentProof = new bytes[](_challenge.stateProof.length - siblingsOffset + 2);
                commitmentProof[0] = _challenge.stateProof[0];
                commitmentProof[1] = _challenge.stateProof[1];
                for (uint256 i = siblingsOffset; i < _challenge.stateProof.length; i++) {
                    commitmentProof[2 + i - siblingsOffset] = _challenge.stateProof[i];
                }

                if (!_verifyConsistencyProof(remoteCommitmentHash, commitmentRoot, commitmentProof)) {
                    return (false, "invalid consistency proof");
                }
            }
        }

        if (targetRollup.getBlockCommitmentHash(_challenge.l2BlockID) == remoteCommitmentHash) {
            return (false, "commitments are consistent");
        }

        return (true, "success");
    }

    function _updateRangeAfterQuestion(
        DataTypes.ChallengeRecord storage _target,
        bool _ack
    ) internal {
        if (_ack) {
            _target.begin = _target.middleBlock.blockID;
            _target.beginBlock = _target.middleBlock;
        } else {
            _target.end = _target.middleBlock.blockID;
            _target.endBlock = _target.middleBlock;
        }

        if (_target.end - _target.begin <= 1) {
            _target.state = 4;
        } else {
            _target.state = 3;
        }
        _target.confirmTime = block.number + WAIT_PERIOD;
    }

    function _resolveChallengeSuccess(string memory _index, string memory _desc) internal returns (uint256, string memory) {
        DataTypes.ChallengeRecord storage target = challengePool[_index];
        IBCPRollupTarget(target.detail.account).rollBackBlock(target.detail.l2BlockID);
        target.state = 5;
        emit ChallengeStateNotify(_index, 5, _desc);
        return (0, "success");
    }

    /* Methods */
    // TODO: create challenge record and check if it is existing
    function ChallengeCreate(DataTypes.ChallengeState memory _challenge, string memory _index, bytes calldata _header) public returns (uint256, string memory) {
        string memory index = _index;

        if (_isChallengeLive(challengePool[index])) {
            emit ChallengeStateNotify(index, challengePool[index].state, "create fail! duplicate challenge");
            return (1, "duplicate challenge");
        }

        (bool legal, string memory desc) = _verifyChallengeRequest(_challenge, _header);
        if (!legal) {
            emit ChallengeStateNotify(index, 0, desc);
            return (2, desc);
        }

        // 保存挑战记录
        challengePool[index].index = index;
        challengePool[index].detail = _challenge;
        challengePool[index].confirmTime = block.number + CONFIRM_PERIOD;
        challengePool[index].challenger = msg.sender;
        challengePool[index].state = 1;
        challengePool[index].begin = chainCheckPoints[_challenge.chainID].blockID;
        challengePool[index].end = _challenge.l1BlockID;

        // 根据 check point 设置 初始区块
        challengePool[index].beginBlock.chainID = _challenge.chainID;
        challengePool[index].beginBlock.blockID = chainCheckPoints[_challenge.chainID].blockID;
        challengePool[index].beginBlock.stateRoot = chainCheckPoints[_challenge.chainID].stateRoot;
        challengePool[index].beginBlock.blockData = chainCheckPoints[_challenge.chainID].blockData;

        // 根据 challenge信息 设置 结束区块
        challengePool[index].endBlock.chainID = _challenge.chainID;
        challengePool[index].endBlock.blockID = _challenge.l1BlockID;
        challengePool[index].endBlock.stateRoot = _challenge.l1StateRoot;
        challengePool[index].endBlock.blockData = _challenge.l1BlockData;

        emit ChallengeStateNotify(index, 1, "create success!");
        return (0, "success");
    }

    // TODO: 1/2 shard to get the challenge record
    function ChallengeQuestion(string memory _index, bool _ack) public returns (uint256, string memory) {
        DataTypes.ChallengeRecord storage target = challengePool[_index];
        if (!_isChallengeLive(target)) {
            emit ChallengeStateNotify(_index, target.state, "fail! challenge not active");
            return (1, "fail! challenge not active");
        }

        if (block.number > target.confirmTime) {
            emit ChallengeStateNotify(_index, target.state, "fail! challenge step expired");
            return (2, "fail! challenge step expired");
        }

        if (target.state == 1) {
            if (_ack) {
                return _resolveChallengeSuccess(_index, "challenge acknowledged");
            }

            target.questioner = msg.sender;
            target.state = 3;
            target.confirmTime = block.number + WAIT_PERIOD;
            emit ChallengeStateNotify(_index, target.state, "success! the commit was questioned");
            return (0, "success");
        }

        if (target.state != 2) {
            emit ChallengeStateNotify(_index, target.state, "fail! not in the correct state!");
            return (3, "fail! not in the correct state!");
        }

        if (msg.sender != target.questioner) {
            emit ChallengeStateNotify(_index, target.state, "fail! illegal questioner!");
            return (4, "fail! illegal questioner!");
        }

        _updateRangeAfterQuestion(target, _ack);
        if (_ack) {
            emit ChallengeStateNotify(_index, target.state, "True! In the last 1/2.");
        } else {
            emit ChallengeStateNotify(_index, target.state, "False! In the before 1/2.");
        }

        return (0, "success");
    }

    //
    // TODO: 1/2 shard to get the challenge record
    function ChallengeResponse(string memory _index, DataTypes.L1BlockInfo memory _middleBlock) public returns (uint256, string memory) {
        DataTypes.ChallengeRecord storage target = challengePool[_index];
        if (target.state != 3) {
            emit ChallengeStateNotify(_index, target.state, "fail! not in the correct state!");
            return (1, "fail! not in the correct state!");
        }

        if (block.number > target.confirmTime) {
            emit ChallengeStateNotify(_index, target.state, "fail! challenge step expired");
            return (2, "fail! challenge step expired");
        }

        if (msg.sender != target.challenger) {
            emit ChallengeStateNotify(_index, target.state, "fail! illegal challenger!");
            return (3, "fail! illegal challenger!");
        }

        uint256 middle = target.begin + ((target.end - target.begin) / 2);
        if (_middleBlock.chainID != target.detail.chainID || _middleBlock.blockID != middle) {
            emit ChallengeStateNotify(_index, target.state, "fail! error block info!");
            return (4, "fail! error block info!");
        }

        target.state = 2;
        target.middleBlock = _middleBlock;
        target.confirmTime = block.number + WAIT_PERIOD;

        emit ChallengeStateNotify(_index, target.state, "challenge response submitted");
        return (0, "success");
    }


    // TODO: verify the final transition
    function FinalChallenge(string memory _index, DataTypes.ZKProof memory _finalProof, bytes calldata _header1, bytes calldata _header2) public returns (uint256, string memory) {
        DataTypes.ChallengeRecord storage target = challengePool[_index];
        if (target.state != 4) {
            emit ChallengeStateNotify(_index, target.state, "fail! not in the final stage!");
            return(1, "fail! incorrect stage!");
        }

        if (block.number > target.confirmTime) {
            emit ChallengeStateNotify(_index, target.state, "fail! challenge step expired");
            return (2, "fail! challenge step expired");
        }

        if (msg.sender != target.challenger) {
            emit ChallengeStateNotify(_index, target.state, "fail! illegal challenger!");
            return (3, "fail! illegal challenger!");
        }

        if (
            !_verifyHeaderBinding(target.beginBlock.blockData, _header1) ||
            !_verifyHeaderBinding(target.endBlock.blockData, _header2)
        ) {
            emit ChallengeStateNotify(_index, target.state, "fail! header proof mismatch");
            return (4, "fail! header proof mismatch");
        }

        (bool hasBeginHeaderRoot, bytes32 beginHeaderRoot,,) = _decodeStateRootBundle(target.beginBlock.stateRoot);
        if (hasBeginHeaderRoot) {
            (bool hasBeginStateRoot, bytes32 beginStateRoot) = HeaderLib.tryStateRoot(_header1);
            if (hasBeginStateRoot && beginStateRoot != beginHeaderRoot) {
                emit ChallengeStateNotify(_index, target.state, "fail! begin header state root mismatch");
                return (5, "fail! begin header state root mismatch");
            }
        }

        (bool hasEndHeaderRoot, bytes32 endHeaderRoot,,) = _decodeStateRootBundle(target.endBlock.stateRoot);
        if (hasEndHeaderRoot) {
            (bool hasEndStateRoot, bytes32 endStateRoot) = HeaderLib.tryStateRoot(_header2);
            if (hasEndStateRoot && endStateRoot != endHeaderRoot) {
                emit ChallengeStateNotify(_index, target.state, "fail! end header state root mismatch");
                return (6, "fail! end header state root mismatch");
            }
        }

        (bool hasParentHash, bytes32 parentHash) = HeaderLib.tryParentHash(_header2);
        if (hasParentHash && parentHash != HeaderLib.canonicalHash(_header1)) {
            emit ChallengeStateNotify(_index, target.state, "fail! header successor mismatch");
            return (7, "fail! header successor mismatch");
        }

        uint256[4] memory proofInputs = _headerHashToProofInputs(HeaderLib.canonicalHash(_header2));

        // 验证状态转换证明，确认状态的合法性
        bool success;
        bytes memory returnData;
        (success, returnData) = address(myZKV).call(
            abi.encodeWithSelector(
                myZKV.verifyProof.selector,
                _finalProof.a, _finalProof.b, _finalProof.c, proofInputs
            )
        );
        bool result;
        (result) = abi.decode(
            (returnData),
            (bool)
        );
        if (! success || ! result) {
            emit ChallengeStateNotify(_index, 0, "create fail! can't verify the proof");
            return (1, "create fail! can't verify the proof");
        }

        return _resolveChallengeSuccess(_index, "final proof success");
    }

    function FinalizeChallenge(string memory _index) public returns (uint256, string memory) {
        DataTypes.ChallengeRecord storage target = challengePool[_index];
        if (!_isChallengeLive(target)) {
            emit ChallengeStateNotify(_index, target.state, "fail! challenge not active");
            return (1, "fail! challenge not active");
        }

        if (block.number <= target.confirmTime) {
            emit ChallengeStateNotify(_index, target.state, "fail! challenge still in dispute window");
            return (2, "fail! challenge still in dispute window");
        }

        if (target.state == 1 || target.state == 2) {
            return _resolveChallengeSuccess(_index, "challenge timeout success");
        }

        target.state = 6;
        emit ChallengeStateNotify(_index, target.state, "challenge timeout failed");
        return (0, "challenge timeout failed");
    }

    // TODO: give a checkPoint of other chain
    function CheckPoint(DataTypes.ChainCheckPoint memory _checkPoint) public onlySuper returns(uint256, string memory) {

        require(_checkPoint.contractAddr != address(0), "invalid rollup address");

        // 验证跨链证明，暂不验证

        // 更新链状态，同时保存旧状态
        string memory index = strConcat(Strings.toString(_checkPoint.chainID), "-");
        index = strConcat(index, Strings.toString(chainCheckPoints[_checkPoint.chainID].index));

        if (chainCheckPoints[_checkPoint.chainID].contractAddr != address(0)) {
            require(
                chainCheckPoints[_checkPoint.chainID].contractAddr == _checkPoint.contractAddr,
                "checkpoint contract mismatch"
            );
            require(
                _checkPoint.blockID > chainCheckPoints[_checkPoint.chainID].blockID,
                "checkpoint block not advanced"
            );
        }

        // 保存旧的状态数据
        checkPointPool[index] =  chainCheckPoints[_checkPoint.chainID];

        // 更新缓存的状态数据
        chainCheckPoints[_checkPoint.chainID].index += 1;
        chainCheckPoints[_checkPoint.chainID].blockID = _checkPoint.blockID;
        chainCheckPoints[_checkPoint.chainID].stateRoot = _checkPoint.stateRoot;
        chainCheckPoints[_checkPoint.chainID].blockData = _checkPoint.blockData;
        chainCheckPoints[_checkPoint.chainID].contractAddr = _checkPoint.contractAddr;

        // notify info
        emit BlockInfoNotify(index, 1, "new info as the chain checkpoint");

        return (0, "success");
    }


    // other functions
    function strConcat(string memory _a, string memory _b) internal pure returns (string memory){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        string memory ret = new string(_ba.length + _bb.length);
        bytes memory bret = bytes(ret);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++) {
            bret[k++] = _ba[i];
        }
        for (uint i = 0; i < _bb.length; i++) {
            bret[k++] = _bb[i];
        }
        return string(ret);
    }
}
