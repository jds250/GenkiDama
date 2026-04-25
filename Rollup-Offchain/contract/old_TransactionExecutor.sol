// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;


/* Internal Imports */
import {DataTypes} from "./DataTypes.sol";
import {MPTVerifier} from "./MPTVerifier.sol";
import {MerkleUtils} from "./MerkleUtils.sol";


contract TransactionExecutor {
    bytes32 constant ZERO_BYTES32 = 0x0000000000000000000000000000000000000000000000000000000000000000;

    MPTVerifier private myMPTV;
    DataTypes private myDT;
    MerkleUtils private myMU;

    /***************
     * Constructor *
     **************/
    constructor(
        address _MPTVerifierAddress,
        address _DataTypesAddress,
        address _MerkleUtilsAddress
    ) {
        myMPTV = MPTVerifier(_MPTVerifierAddress);
        myDT = DataTypes(_DataTypesAddress);
        myMU = MerkleUtils(_MerkleUtilsAddress);
    }

    function VerifyTransactions(
        DataTypes.L2Transaction[] memory _transactions,
        DataTypes.Block memory _targetBlock,
        MPTVerifier.MerkleProof[] memory _stateProof
    ) public view returns (bool, string memory) {
        // get Transition
        DataTypes.L2Transition[] memory L2Transitions = GetTransition(_transactions);

        // get LocalTransaction
        DataTypes.LocalTransition[] memory localTransitions = GetLocalTransition(_transactions, _targetBlock);
        bytes[] memory ltBytes = LocalTransitionsEncode(localTransitions);

        // verify local Transaction
        bytes32 tmpRoot = myMU.getMerkleRoot(ltBytes);
        if (_targetBlock.localTransitionRoot != tmpRoot) {
            return (false, "targetBlock.localStateRoot is Wrong");
        }

        // verify state change (with l2Transition and state proof)
        for (uint256 i = 0; i < localTransitions.length; i ++) {
            uint256 res = VerifyOneTransition(L2Transitions, _stateProof, i);
            if (res == 1) {
                return (false, "Proof fail: incorrect value!");
            } else if (res == 2) {
                return (false, "Proof fail: invalid proof!");
            }
        }

        return (true, "Proof Success!");
    }


    // VerifyOneTransition 验证单个转换的正确性
    // returns: 0-true; 1-incorrect value; 2-invalid proof
    function VerifyOneTransition(
        DataTypes.L2Transition[] memory _transition,
        MPTVerifier.MerkleProof[] memory _stateProof,
        uint256 index
    ) public view returns(uint256) {
        uint256 first = bytesToUint(_stateProof[2*index].expectedValue);
        uint256 second = bytesToUint(_stateProof[2*index+1].expectedValue);

        if (_transition[index].style == 1) {
            if ((second - first) != _transition[index].value) {
                return 1;
            }
        }

        if (_transition[index].style == 2) {
            if ((first - second) != _transition[index].value) {
                return 1;
            }
        }

        if (!myMPTV.verifyProof(_stateProof[2*index]) || !myMPTV.verifyProof(_stateProof[2*index+1])) {
            return 2;
        }

        return 0;
    }

    // GetTransition Translate Transaction to L2 Transition
    function GetTransition(DataTypes.L2Transaction[] memory _transactions) public pure returns (DataTypes.L2Transition[] memory) {
        uint256 cnt = 0;
        for (uint256 i = 0; i < _transactions.length; i ++) {
            if (_transactions[i].style == 0) {
                cnt += 2;
            }
            if (_transactions[i].style == 1 || _transactions[i].style == 2) {
                cnt += 1;
            }
        }

        DataTypes.L2Transition[] memory resTransit = new DataTypes.L2Transition[](cnt);

        uint256 index = 0;
        for (uint256 i = 0; i < _transactions.length; i ++) {
            // 添加输入状态转化
            if (_transactions[i].style != 2) {
                resTransit[index].account = _transactions[i].fromAddr;
                resTransit[index].value = _transactions[i].value;
                resTransit[index].style = 2;
                index ++;
            }

            // 添加输出状态转化
            if (_transactions[i].style != 1) {
                resTransit[index].account = _transactions[i].toAddr;
                resTransit[index].value = _transactions[i].value;
                resTransit[index].style = 1;
                index ++;
            }
        }

        return resTransit;
    }

    // GetTransition Translate Transaction to L2 Transition
    function GetLocalTransition(DataTypes.L2Transaction[] memory _transactions, DataTypes.Block memory _targetBlock) public pure returns (DataTypes.LocalTransition[] memory) {
        uint256 cnt = 0;
        for (uint256 i = 0; i < _transactions.length; i ++) {
            if (_transactions[i].style == 1 || _transactions[i].style == 2) {
                cnt += 1;
            }
        }

        DataTypes.LocalTransition[] memory resTransit = new DataTypes.LocalTransition[](cnt);

        uint256 index = 0;
        for (uint256 i = 0; i < _transactions.length; i ++) {
            // 添加输入状态转化 L1-L2
            if (_transactions[i].style == 2) {
                resTransit[index].account = _transactions[i].fromAddr;
                resTransit[index].value = _transactions[i].value;
                resTransit[index].chainID = _transactions[i].chainID;
                resTransit[index].blockNum = _targetBlock.height;
                resTransit[index].style = 2;
                index ++;
            }

            // 添加输出状态转化 L2-L1
            if (_transactions[i].style == 1) {
                resTransit[index].account = _transactions[i].toAddr;
                resTransit[index].value = _transactions[i].value;
                resTransit[index].chainID = _transactions[i].chainID;
                resTransit[index].blockNum = _targetBlock.height;
                resTransit[index].style = 1;
                index ++;
            }
        }

        return resTransit;
    }

    // GetAllTransition Of L1
    function GetAllTransition(DataTypes.L2Transaction[] memory _transactions, DataTypes.Block memory _targetBlock) public pure returns (DataTypes.LocalTransition[] memory) {
        uint256 cnt = _transactions.length;

        DataTypes.LocalTransition[] memory resTransit = new DataTypes.LocalTransition[](cnt);

        uint256 index = 0;
        for (uint256 i = 0; i < _transactions.length; i ++) {
            // 添加输入状态转化 L1-L2
            if (_transactions[i].style == 2) {
                resTransit[index].account = _transactions[i].fromAddr;
                resTransit[index].value = _transactions[i].value;
                // resTransit[index].chainID = _transactions[i].chainID;
                resTransit[index].chainID = 1;
                resTransit[index].blockNum = _targetBlock.height;
                resTransit[index].style = 2;
                index ++;
            }

            // 添加输出状态转化 L2-L1
            if (_transactions[i].style == 1) {
                resTransit[index].account = _transactions[i].toAddr;
                resTransit[index].value = _transactions[i].value;
//                resTransit[index].chainID = _transactions[i].chainID;
                resTransit[index].chainID = 1;
                resTransit[index].blockNum = _targetBlock.height;
                resTransit[index].style = 1;
                index ++;
            }
        }

        return resTransit;
    }

    // LocalTransitionsEncode encode Transition List
    function LocalTransitionsEncode(DataTypes.LocalTransition[] memory _transitions) public view returns(bytes[] memory) {
        bytes[] memory res = new bytes[](_transitions.length);

        for (uint256 i = 0; i < _transitions.length; i ++) {
            res[i] = myDT.LocalTransitionToBytes(_transitions[i]);
        }

        return res;
    }

    // bytesToUint
    // function bytesToUint(bytes memory b) public returns (uint256){
    //     uint256 number;
    //     for(uint256 i=0;i<b.length;i++){
    //         number = number + int(b[i])*(2**(8*(b.length-(i+1))));
    //     }
    //     return number;
    // }

    function bytesToUint(bytes memory b) public pure returns (uint256){
        uint256 number;
        for(uint i=0;i<b.length;i++){
            number = number + uint8(b[i]);
        }
        return number;
    }

}
