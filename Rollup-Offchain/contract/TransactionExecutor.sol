// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;


/* Internal Imports */
import {DataTypes} from "./DataTypes.sol";
import {MPTVerifier} from "./MPTVerifier.sol";
import {MerkleUtils} from "./MerkleUtils.sol";
import {LocalStateManager} from "./LocalStateManager.sol";


contract TransactionExecutor {
    bytes32 constant ZERO_BYTES32 = 0x0000000000000000000000000000000000000000000000000000000000000000;

    MPTVerifier private myMPTV;
    DataTypes private myDT;
    MerkleUtils private myMU;
    // Local State Manager
    LocalStateManager myLSM;
    uint256 public sourceChainID;
    uint256 public targetChainID;

    /***************
     * Constructor *
     **************/
    constructor(
        address _MPTVerifierAddress,
        address _DataTypesAddress,
        address _MerkleUtilsAddress,
        address _LocalStateManagerAddress,
        uint256 _sourceChainID,
        uint256 _targetChainID
    ) {
        myMPTV = MPTVerifier(_MPTVerifierAddress);
        myDT = DataTypes(_DataTypesAddress);
        myMU = MerkleUtils(_MerkleUtilsAddress);
        myLSM = LocalStateManager(_LocalStateManagerAddress);
        sourceChainID = _sourceChainID;
        targetChainID = _targetChainID;
    }

    function _validateTransaction(DataTypes.L2Transaction memory _transaction) internal view returns (bool, string memory) {
        if (bytes(_transaction.fromAddr).length == 0 || bytes(_transaction.toAddr).length == 0) {
            return (false, "empty account");
        }

        if (_transaction.value == 0) {
            return (false, "zero value transaction");
        }

        if (_transaction.style == 1) {
            if (_transaction.chainID != sourceChainID) {
                return (false, "unexpected source chain id");
            }
            return (true, "");
        }

        if (_transaction.style == 2) {
            if (_transaction.chainID != targetChainID) {
                return (false, "unexpected source chain id");
            }
            return (true, "");
        }

        return (false, "unsupported transaction style");
    }

    function _resolveTransitionChains(
        DataTypes.L2Transaction memory _transaction
    ) internal view returns (uint256 debitChainID, uint256 creditChainID) {
        if (_transaction.style == 1) {
            return (sourceChainID, targetChainID);
        }

        if (_transaction.style == 2) {
            return (targetChainID, sourceChainID);
        }

        revert("unsupported transaction style");
    }

    function VerifyTransactions(
        DataTypes.L2Transaction[] memory _transactions,
        bytes32 _tsRoot
    ) public view returns (bool, string memory) {
        for (uint256 i = 0; i < _transactions.length; i++) {
            (bool txValid, string memory txDesc) = _validateTransaction(_transactions[i]);
            if (!txValid) {
                return (false, txDesc);
            }
        }

        // get L2Transaction
        DataTypes.L2Transition[] memory l2Transitions = GetL2Transitions(_transactions);

        // verify local Transition with root
        bytes32 tmpRoot = myDT.L2TransitionListHash(l2Transitions);
        if (_tsRoot != tmpRoot) {
            return (false, "targetBlock.localStateRoot is Wrong");
        }

        // verify the local transition with local state
        bool result = myLSM.verifyLocalState(l2Transitions);

        // lock local state
        // bool success;
        // bytes memory returnData;
        // (success, returnData) = address(myLSM).call(
        //     abi.encodeWithSelector(
        //         myLSM
        //         .verifyLocalState
        //         .selector,
        //         l2Transitions
        //     )
        // );

        // bool result = false;
        // if (success) {
        //     (result) = abi.decode(
        //         (returnData),
        //         (bool)
        //     );
        // }

        if (result) {
            return (true, "Proof Success!");
        } else {
            return (false, "no enough token");
        }
    }


    // GetL2Transitions Translate Transaction to L2 Transition
    function GetL2Transitions(DataTypes.L2Transaction[] memory _transactions) public view returns (DataTypes.L2Transition[] memory) {
        uint256 cnt = 2 * _transactions.length;

        DataTypes.L2Transition[] memory resTransit = new DataTypes.L2Transition[](cnt);

        uint256 index = 0;
        for (uint256 i = 0; i < _transactions.length; i ++) {
            (bool txValid,) = _validateTransaction(_transactions[i]);
            require(txValid, "invalid transaction");

            (uint256 debitChainID, uint256 creditChainID) = _resolveTransitionChains(_transactions[i]);
            resTransit[index].chainID = debitChainID;
            resTransit[index + 1].chainID = creditChainID;

            // 添加输入状态转化
            resTransit[index].account = _transactions[i].fromAddr;
            resTransit[index].value = _transactions[i].value;
            resTransit[index].style = 1;

            // 添加输出状态转化
            resTransit[index + 1].account = _transactions[i].toAddr;
            resTransit[index + 1].value = _transactions[i].value;
            resTransit[index + 1].style = 2;

            index += 2;
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


    function bytesToUint(bytes memory b) public pure returns (uint256){
        uint256 number;
        for(uint i=0;i<b.length;i++){
            number = number + uint8(b[i]);
        }
        return number;
    }

}
