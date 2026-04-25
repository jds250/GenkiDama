// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

/* Internal Imports */
import {DataTypes} from "./DataTypes.sol";
import {IZKVerifier} from "./IZKVerifier.sol";

interface IChallengeRollupTarget {
    function isBlockInDispute(uint256 _blockNumber) external view returns (bool);
    function getBlockCommitmentHash(uint256 _blockNumber) external view returns (bytes32);
    function rollBackBlock(uint256 _targetBlockHeight) external;
}

/// @title 挑战状态管理
// 专门用于管理挑战状态
contract ChallengeManager {
    /*** Fields ***/
    /* Contract Instance */
    // The Data Type library
    DataTypes myDT;
    // The Data Type library
    IZKVerifier myZKV;

    /* Const */
    bytes32 public constant ZERO_BYTES32 = 0x0000000000000000000000000000000000000000000000000000000000000000;
    // TODO: Set a reasonable wait period
    uint256 constant WAIT_PERIOD = 4;
    uint256 constant CONFIRM_PERIOD = 10;

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
        address _zkVerifierAddress
    ) {
        chainID = _chainID;
        myDT = DataTypes(_dataTypesAddress);
        myZKV = IZKVerifier(_zkVerifierAddress);
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

    function _getRemoteCommitmentHash(DataTypes.ChallengeState memory _challenge) internal pure returns (bytes32, bool) {
        if (_challenge.stateProof.length == 0 || _challenge.stateProof[0].length != 32) {
            return (ZERO_BYTES32, false);
        }

        return (_decodeBytes32(_challenge.stateProof[0]), true);
    }

    function _headerHashToProofInputs(bytes32 _headerHash) internal pure returns (uint256[4] memory inputs) {
        uint256 value = uint256(_headerHash);
        uint256 mask = uint256(type(uint64).max);

        inputs[0] = value & mask;
        inputs[1] = (value >> 64) & mask;
        inputs[2] = (value >> 128) & mask;
        inputs[3] = (value >> 192) & mask;
    }

    function _verifyChallengeRequest(
        DataTypes.ChallengeState memory _challenge
    ) internal view returns (bool, string memory) {
        if (_challenge.account == address(0)) {
            return (false, "missing rollup address");
        }

        IChallengeRollupTarget targetRollup = IChallengeRollupTarget(_challenge.account);
        if (!targetRollup.isBlockInDispute(_challenge.l2BlockID)) {
            return (false, "challenge window closed");
        }

        (bytes32 remoteCommitmentHash, bool ok) = _getRemoteCommitmentHash(_challenge);
        if (!ok) {
            return (false, "missing commitment proof");
        }

        (, , bool hasCommitmentRoot, bytes32 commitmentRoot) = _decodeStateRootBundle(_challenge.l1StateRoot);
        if (hasCommitmentRoot && !_verifyConsistencyProof(remoteCommitmentHash, commitmentRoot, _challenge.stateProof)) {
            return (false, "invalid consistency proof");
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
        IChallengeRollupTarget(target.detail.account).rollBackBlock(target.detail.l2BlockID);
        target.state = 5;
        emit ChallengeStateNotify(_index, 5, _desc);
        return (0, "success");
    }

    /* Methods */
    // TODO: create challenge record and check if it is existing
    function ChallengeCreate(DataTypes.ChallengeState memory _challenge, string memory _index, DataTypes.ZKProof memory _proof) public returns (uint256, string memory) {
        string memory index = _index;

        if (_isChallengeLive(challengePool[index])) {
            emit ChallengeStateNotify(index, challengePool[index].state, "create fail! duplicate challenge");
            return (1, "duplicate challenge");
        }

        if (_challenge.l1BlockData.length != 32) {
            emit ChallengeStateNotify(index, 0, "create fail! invalid header hash");
            return (2, "create fail! invalid header hash");
        }

        // 验证挑战数据的合法性 zk verify
        bool success;
        bytes memory returnData;
        (success, returnData) = address(myZKV).call(
            abi.encodeWithSelector(
                myZKV.verifyProof.selector,
                _proof.a, _proof.b, _proof.c, _headerHashToProofInputs(_decodeBytes32(_challenge.l1BlockData))
            )
        );
        if (! success) {
            emit ChallengeStateNotify(index, 0, "create fail!  can't verify the proof");
            return (1, "create fail! can't verify the proof");
        }
        bool result;
        (result) = abi.decode(
            (returnData),
            (bool)
        );

        if (!result) {
            emit ChallengeStateNotify(index, 0, "create fail!  illegal zk proof");

            return (3, "create fail! illegal zk proof");
        }

        (bool legal, string memory desc) = _verifyChallengeRequest(_challenge);
        if (!legal) {
            emit ChallengeStateNotify(index, 0, desc);
            return (4, desc);
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
    function FinalChallenge(string memory _index, DataTypes.ZKProof memory _finalProof) public returns (uint256, string memory) {
        // 只有在挑战处于最终状态时需要使用
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

        if (target.endBlock.blockData.length != 32) {
            emit ChallengeStateNotify(_index, target.state, "fail! invalid header hash");
            return (4, "fail! invalid header hash");
        }

        // 验证状态转换证明，确认状态的合法性
        bool success;
        bytes memory returnData;
        (success, returnData) = address(myZKV).call(
            abi.encodeWithSelector(
                myZKV.verifyProof.selector,
                _finalProof.a, _finalProof.b, _finalProof.c, _headerHashToProofInputs(_decodeBytes32(target.endBlock.blockData))
            )
        );
        bool result;
        (result) = abi.decode(
            (returnData),
            (bool)
        );
        if (! success || ! result) {
            emit ChallengeStateNotify(_index, 0, "create fail! can't verify the proof");
            return (5, "create fail! can't verify the proof");
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
    function CheckPoint(DataTypes.ChainCheckPoint memory _checkPoint) public returns(uint256, string memory) {

        // 验证跨链证明，暂不验证

        // 更新链状态，同时保存旧状态
        string memory index = strConcat(Strings.toString(_checkPoint.chainID), "-");
        index = strConcat(index, Strings.toString(chainCheckPoints[_checkPoint.chainID].index));

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
