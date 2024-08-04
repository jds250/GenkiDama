// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

/* Internal Imports */
import {DataTypes} from "./DataTypes.sol";
import {ZKVerifier} from "./ZKVerifier.sol";

/// @title 挑战状态管理
// 专门用于管理挑战状态
contract BCPManager {
    /*** Fields ***/
    /* Contract Instance */
    // The Data Type library
    DataTypes myDT;
    // The Data Type library
    ZKVerifier myZKV;

    /* Const */
    bytes32 public constant ZERO_BYTES32 = 0x0000000000000000000000000000000000000000000000000000000000000000;
    // TODO: Set a reasonable wait period
    uint256 constant WAIT_PERIOD = 4;
    uint256 constant CONFIRM_PERIOD = 10;

    /* rollup system state */
    // All the blocks!
    uint256 public chainID;
    mapping(string=>DataTypes.ChallengeRecord) public challengePool;

    // checkpoint record
    mapping(string=>DataTypes.ChainCheckPoint) public checkPointPool;
    DataTypes.ChainCheckPoint[100] public chainCheckPoints;

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
        myZKV = ZKVerifier(_zkVerifierAddress);
        superAddress = msg.sender;
    }

    modifier onlySuper() {
        require(
            msg.sender == address(superAddress),
            "Only super admin may perform action"
        );
        _;
    }

    /* Methods */
    // TODO: create challenge record and check if it is existing
    function ChallengeCreate(DataTypes.ChallengeState memory _challenge, string memory _index, bytes calldata _header) public returns (uint256, string memory) {
        // 确认挑战没有重复
        // string memory index = strConcat(Strings.toString(_challenge.chainID), "-");
        // index = strConcat(index, Strings.toString(_challenge.l1BlockID));
        // index = strConcat(index, "-");
        // index = strConcat(index, Strings.toString(_challenge.l2BlockID));
        string memory index = _index;

        // 验证header 哈希值
        bytes32 header_hash = keccak256(_header);
        if (header_hash != header_hash) {
            return (1, "fail");
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
        DataTypes.ChallengeRecord memory target = challengePool[_index];
        string memory descStr;

        // 判断挑战者是否正确
        if (target.state != 1 && target.state != 2) {
            descStr = "fail! not in the correct state!";
            // emit ChallengeStateNotify(_index, target.state, "fail! not in the correct state!");
        } else if (target.state == 1) { // 如果未被挑战
            if (! _ack) {
                challengePool[_index].questioner = msg.sender;
                challengePool[_index].state = 3; // wait for challenger response
                challengePool[_index].confirmTime = block.number + WAIT_PERIOD;

                descStr = "success! the commit was challenged!";
            }
        } else if (target.state == 2) { // 如果待质疑
            // 确定是质疑者本人才能发起挑战
            if (msg.sender == target.questioner) {
                if (!_ack) {
                    // 如果否认挑战者提供的区块，争议区间在前 1/2
                    challengePool[_index].end = (target.begin + target.end)/2;
                    challengePool[_index].endBlock = challengePool[_index].middleBlock;
                } else {
                    // 如果承认挑战者提供的区块，争议区间在后 1/2
                    challengePool[_index].begin = (target.begin + target.end)/2;
                    challengePool[_index].beginBlock = challengePool[_index].middleBlock;
                }

                // 确认是否已经无法二分
                if (target.end - target.begin <= 1) {
                    challengePool[_index].state = 4; // 设置为待最终证明
                } else {
                    challengePool[_index].state = 3; // 设置为待回复
                }

                // 根据最新结果发出事件通知
                if (_ack) {
                    descStr = "True! In the last 1/2.";
                } else {
                    descStr = "False! In the before 1/2.";
                }

                // 成功更新后旧延长争议时间
                challengePool[_index].confirmTime = block.number + WAIT_PERIOD;
            }
        }

        emit ChallengeStateNotify(_index, challengePool[_index].state, descStr);

        return (0, "nothing");
    }

    //
    // TODO: 1/2 shard to get the challenge record
    function ChallengeResponse(string memory _index, DataTypes.L1BlockInfo memory _middleBlock) public returns (uint256, string memory) {
        DataTypes.ChallengeRecord memory target = challengePool[_index];
        uint256 middle = (target.end - target.begin)/2;

        // 判断挑战者是否正确
        string memory descStr;
        if (target.state != 3) {
            descStr = "fail! not in the correct state!";
        } else if (target.state == 3) { // 如果待回应 (只有这里可以挑战)
            if (msg.sender != target.challenger) {
                descStr = "fail! illegal challenger!";
            }

            if (_middleBlock.chainID != target.detail.chainID || _middleBlock.blockID != middle) {
                descStr = "fail! error block info!";
            }

            // 如果确定身份合法 且 区块信息正确，进入下一级挑战
            challengePool[_index].state = 2;
            challengePool[_index].middleBlock = _middleBlock;

            // 发起通知
            descStr = string(_middleBlock.blockData);
        }
        emit ChallengeStateNotify(_index, challengePool[_index].state, descStr);

        return (0, "nothing");
    }


    // TODO: verify the final transition
    function FinalChallenge(string memory _index, DataTypes.ZKProof memory _finalProof, bytes calldata _header1, bytes calldata _header2) public returns (uint256, string memory) {
        // 只有在挑战处于最终状态时需要使用
        DataTypes.ChallengeRecord memory target = challengePool[_index];
        if (target.state != 4) {
            emit ChallengeStateNotify(_index, target.state, "fail! not in the final stage!");
            return(1, "fail! incorrect stage!");
        }

        // TODO: 确认证明是针对目标数据的
        // 暂时无法实现

        // 验证状态转换证明，确认状态的合法性
        bool success;
        bytes memory returnData;
        (success, returnData) = address(myZKV).call(
            abi.encodeWithSelector(
                myZKV.verifyProof.selector,
                _finalProof.a, _finalProof.b, _finalProof.c
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


        return (0, "success");
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