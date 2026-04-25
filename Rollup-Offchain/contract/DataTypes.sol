// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

contract DataTypes {

    // ***************************
    //        Data Struct
    // ***************************
    struct Block {
        uint256 height;
        bytes32 txRoot; // store root of the transaction batch
        bytes32 localTransitionRoot; // state root of all local state

        bytes32 stateRoot;
        uint256 blockSize; /* the info to commit and store */
        uint256 commitTime;
    }

    // assuming only two blockchains
    struct L2Transaction {
        string fromAddr;
        string toAddr;
        uint256 value;
        uint256 chainID;
        uint256 style; //  1-L1 to L2; 2-L2 to L1
    }

    struct LocalState {
        string account;
        uint256 value; // value is usable
        uint256 lock; // value be used
    }

    struct LocalTransition {
        uint256 chainID;
        uint256 blockNum; // block index
        string account;
        uint256 value; // value is usable
        uint256 style; // style = deposit refund lock-from-l1 lock-from-l2 unlock-to-l1 unlock-to-l2
        uint256 unlockTime; // the block high to apply
    }

    struct L2Transition {
        string account;
        uint256 value; // value is usable
        uint256 style; // 1-ACCOUNT-sub; 2-ACCOUNT-add
        uint256 chainID;
    }

    struct ChallengeState {
        uint256 chainID;
        uint256 l2BlockID;
        uint256 l1BlockID;
        bytes l1StateRoot;
        bytes l1BlockData;
        bytes[] stateProof;
        address account;
    }

    struct ChallengeRecord {
        string index; // chainID-L2ID
        ChallengeState detail;
        uint256 confirmTime;
        address challenger;
        address questioner;
        uint256 state; // 0-不存在，1-未被质疑，2-待质疑，3-待回应，4-待证明（最终），5-质疑成功 6-挑战成功
        // 交互式挑战需要的区块信息
        uint256 begin; // 该链的上一个区块记录
        uint256 end; // 该链当前的区块记录
        L1BlockInfo beginBlock;
        L1BlockInfo endBlock;
        L1BlockInfo middleBlock;
    }

    struct ZKProof {
        uint256[2] a;
        uint256[2][2] b;
        uint256[2] c;
    }

    struct ChainCheckPoint {
        uint256 index;
        uint256 chainID;
        uint256 chainType;
        uint256 blockID;
        bytes stateRoot;
        bytes blockData;
        address contractAddr;
    }

    struct L1BlockInfo {
        uint256 chainID;
        uint256 blockID;
        bytes stateRoot;
        bytes blockData;
    }

    // ***************************
    //      toHash function
    // ***************************
    // L2Transition
    function L2TransitionToBytes(L2Transition memory _l2t) public pure returns(bytes memory) {
        return  abi.encodePacked(
            _l2t.account,
            _l2t.value,
            _l2t.style,
            _l2t.chainID
        );
    }

    function BytesToL2Transition(bytes memory _ltBytes) public pure returns(L2Transition memory) {
        (
            uint256 _value,
            string memory _account,
            uint256 _style,
            uint256 _chainID
        ) = abi.decode(
            (_ltBytes),
            (uint256, string, uint256, uint256)
        );

        L2Transition memory res = L2Transition(_account, _value, _style, _chainID);

        return res;
    }

    // LocalTransition
    function LocalTransitionToBytes(LocalTransition memory _lt) public pure returns(bytes memory) {
        return  abi.encodePacked(
            _lt.chainID,
            _lt.blockNum,
            _lt.account,
            _lt.value,
            _lt.style,
            _lt.unlockTime
        );
    }

    function BytesToLocalTransition(bytes memory _ltBytes) public pure returns(LocalTransition memory) {
        (
            uint256 _chainID,
            uint256 _blockNum,
            string memory _account,
            uint256 _value,
            uint256 _style,
            uint256 _unlockTime
        ) = abi.decode(
            (_ltBytes),
            (uint256, uint256, string, uint256, uint256, uint256)
        );

        LocalTransition memory res = LocalTransition(_chainID, _blockNum, _account, _value, _style, _unlockTime);

        return res;
    }

    // LocalState
    function LocalStateToBytes(LocalState memory _ls) public pure returns(bytes memory) {
        return  abi.encodePacked(
            _ls.account,
            _ls.value,
            _ls.lock
        );
    }

    function BytesToLocalState(bytes memory _ltBytes) public pure returns(LocalState memory) {
        (
            string memory _account,
            uint256 _value,
            uint256 _lock
        ) = abi.decode(
            (_ltBytes),
            (string, uint256, uint256)
        );

        LocalState memory res = LocalState( _account, _value, _lock);

        return res;
    }

    // L2Transaction
    function L2TransactionToBytes(L2Transaction memory _tx) public pure returns(bytes memory) {
        return  abi.encodePacked(
            _tx.fromAddr,
            _tx.toAddr,
            _tx.value,
            _tx.chainID,
            _tx.style
        );
    }

    function BytesToL2Transaction(bytes memory _txBytes) public pure returns(L2Transaction memory) {
        (
            string memory fromAddr,
            string memory toAddr,
            uint256 value,
            uint256 chainID,
            uint256 style
        ) = abi.decode(
            (_txBytes),
            (string, string, uint256, uint256, uint256)
        );

        L2Transaction memory res = L2Transaction( fromAddr, toAddr, value, chainID, style);

        return res;
    }

    // Hash List
    function L2TransitionListHash(L2Transition[] memory _transitions) public view returns(bytes32) {
        bytes[] memory res = new bytes[](_transitions.length);

        uint256 fullLen = 0;
        for (uint256 i = 0; i < _transitions.length; i ++) {
            res[i] = L2TransitionToBytes(_transitions[i]);
            fullLen += res[i].length;
        }

        // get full len
        bytes memory fullRes = new bytes(fullLen);
        uint256 index = 0;
        for (uint256 i = 0; i < res.length; i ++) {
            for (uint256 j = 0; j < res[i].length; j ++) {
                fullRes[index] = res[i][j];
                index ++;
            }
        }

        return keccak256(fullRes);
    }


    function TransactionsListHash(L2Transaction[] memory _transactions) public view returns(bytes32) {
        bytes[] memory res = new bytes[](_transactions.length);

        uint256 fullLen = 0;
        for (uint256 i = 0; i < _transactions.length; i ++) {
            res[i] = L2TransactionToBytes(_transactions[i]);
            fullLen += res[i].length;
        }

        // get full len
        bytes memory fullRes = new bytes(fullLen);
        uint256 index = 0;
        for (uint256 i = 0; i < res.length; i ++) {
            for (uint256 j = 0; j < res[i].length; j ++) {
                fullRes[index] = res[i][j];
                index ++;
            }
        }


        return keccak256(fullRes);
    }


    // LocalTransitionsEncode encode LocalTransition List
    function L2TransitionsEncode(DataTypes.LocalTransition[] calldata _transitions) public view returns(bytes[] memory) {
        bytes[] memory res = new bytes[](_transitions.length);

        for (uint256 i = 0; i < _transitions.length; i ++) {
            DataTypes.LocalTransition memory ts = _transitions[i];
            res[i] = LocalTransitionToBytes(ts);
        }

        return res;
    }

    // LocalTransitionsEncode encode LocalTransition List
    function LocalTransitionsEncode(DataTypes.LocalTransition[] calldata _transitions) public view returns(bytes[] memory) {
        bytes[] memory res = new bytes[](_transitions.length);

        for (uint256 i = 0; i < _transitions.length; i ++) {
            DataTypes.LocalTransition memory ts = _transitions[i];
            res[i] = LocalTransitionToBytes(ts);
        }

        return res;
    }

    function L2TransactionsEncode(DataTypes.L2Transaction[] memory _transactions) public view returns(bytes[] memory) {
        bytes[] memory res = new bytes[](_transactions.length);

        for (uint256 i = 0; i < _transactions.length; i ++) {
            res[i] = L2TransactionToBytes(_transactions[i]);
        }

        return res;
    }

    // address to string
    // function toString(address account) public pure returns(string memory) {
    //     return Strings.toHexString(uint160(address), 20);
    // }
}
