// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;

/* Internal Imports */
import {DataTypes as dt} from "./DataTypes.sol";
import {MerkleUtils} from "./MerkleUtils.sol";


/// @title 本地状态管理
contract LocalStateManager {

    /*** Fields ***/
    /* Contract Instance */
    // The Merkle Tree library (currently a contract for ease of testing)
    MerkleUtils merkleUtils;
    // The Data Type library
    dt dataTypes;
    /* Const */
    bytes32 public constant ZERO_BYTES32 = 0x0000000000000000000000000000000000000000000000000000000000000000;
    // State tree height
    uint256 constant STATE_TREE_HEIGHT = 32;
    // TODO: Set a reasonable wait period
    uint256 constant WITHDRAW_WAIT_PERIOD = 4;

    /* rollup system state */
    // All the blocks!
    uint256 public chainID;
    mapping(string=>dt.LocalState) public localStatePool; // the coin to transfer in/out layer2
    mapping(uint256=>dt.LocalTransition) public lockTransitionPool;
    uint256 public lockTransitionNum;

    // authority
    address public committerAddress;
    address public superAddress;
    address public rollupAddress;

    /* Events */
    event ErrorNotify(uint256 state, string error);
    event LocalStateNotify(string account, string style, uint256 value); // style = deposit refund lock-from-l1 lock-from-l2 unlock-to-l1 unlock-to-l2
    event ConfirmState(uint256 unlock, uint256 now, string result);


    /***************
     * Constructor *
     **************/
    constructor(
        uint256 _chainID,
        address _merkleUtilsAddress,
        address _dataTypesAddress
    ) {
        chainID = _chainID;
        merkleUtils = MerkleUtils(_merkleUtilsAddress);
        dataTypes = dt(_dataTypesAddress);
        superAddress = msg.sender;
        lockTransitionNum = 0;
    }

    modifier onlySuper() {
        require(
            msg.sender == address(superAddress),
            "Only super admin may perform action"
        );
        _;
    }

    modifier onlyRollup() {
        require(
            msg.sender == address(rollupAddress),
            "Only rollup contract may perform action"
        );
        _;
    }

    function setRollupAddress(address _rollupAddress) external onlySuper {
        require(_rollupAddress != address(0), "rollup address is zero");
        rollupAddress = _rollupAddress;
    }

    function _accumulatedDebitExceedsBalance(
        dt.L2Transition[] memory _transits
    ) internal view returns (bool) {
        string[] memory touchedAccounts = new string[](_transits.length);
        uint256[] memory accumulatedDebit = new uint256[](_transits.length);
        uint256 touchedCount = 0;

        for (uint256 i = 0; i < _transits.length; i++) {
            if (_transits[i].chainID != chainID || _transits[i].style != 1) {
                continue;
            }

            bytes32 accountHash = keccak256(bytes(_transits[i].account));
            uint256 accountIndex = touchedCount;
            bool found = false;
            for (uint256 j = 0; j < touchedCount; j++) {
                if (keccak256(bytes(touchedAccounts[j])) == accountHash) {
                    accountIndex = j;
                    found = true;
                    break;
                }
            }

            if (!found) {
                touchedAccounts[touchedCount] = _transits[i].account;
                accountIndex = touchedCount;
                touchedCount ++;
            }

            accumulatedDebit[accountIndex] += _transits[i].value;
            if (accumulatedDebit[accountIndex] > localStatePool[_transits[i].account].value) {
                return true;
            }
        }

        return false;
    }

    /* Methods */
    // depositCoin deposit and generate a localStateRecord
    function depositCoin(string memory _account, uint256 _value) public onlySuper {
        if (bytes(_account).length == 0) {
            emit ErrorNotify(1, "account string can't be empty");
        }

        if (bytes(localStatePool[_account].account).length == 0) {
            localStatePool[_account].account = _account;
            localStatePool[_account].value = 0;
            localStatePool[_account].lock = 0;
        }

        localStatePool[_account].value += _value;

        emit LocalStateNotify(_account, "deposit", _value);
    }

    // refundCoin: refund the coin from localStatePool
    function refundCoin(string memory _account, uint256 _value) public onlySuper {
        if (bytes(localStatePool[_account].account).length == 0) {
            emit ErrorNotify(1, "account string can't be empty");
        }

        localStatePool[_account].value -= _value;
        emit LocalStateNotify(_account, "refund", _value);
    }

    // get account balance
    function getLocalState(string memory _account) public view returns(dt.LocalState memory) {
        return localStatePool[_account];
    }

    // verifyAndLockLocalState if all relative the local state exist, lock them. Otherwise, refuse the local state
    function verifyAndLockLocalState(dt.L2Transition[] memory _transits, uint256 _blockNum, uint256 _lockTime) public onlyRollup returns (bool) {
        dt.LocalTransition[] memory tmpTransitionList = new dt.LocalTransition[](
            _transits.length
        );
        uint256 tmpNum = 0;

        if (_accumulatedDebitExceedsBalance(_transits)) {
            emit ErrorNotify(1, "illegal transition: cumulative account balance is not enough");
            return false;
        }

        // decode all LocalState and verify
        for (uint256 i = 0; i < _transits.length; i++) {
            if (_transits[i].chainID == chainID) {
                tmpTransitionList[tmpNum].chainID = _transits[i].chainID;
                tmpTransitionList[tmpNum].blockNum = _blockNum;
                tmpTransitionList[tmpNum].account = _transits[i].account;
                tmpTransitionList[tmpNum].value = _transits[i].value;
                tmpTransitionList[tmpNum].style = _transits[i].style;
                tmpTransitionList[tmpNum].unlockTime = _lockTime;

                tmpNum ++;
            }
        }

        // record the local transition and lock the corresponding fund
        for (uint256 i = 0; i < tmpNum; i++) {
            if (tmpTransitionList[i].style == 1) {
                localStatePool[tmpTransitionList[i].account].value -= tmpTransitionList[i].value;
                localStatePool[tmpTransitionList[i].account].lock  += tmpTransitionList[i].value;
                emit LocalStateNotify(tmpTransitionList[i].account, "lock-from-l1", tmpTransitionList[i].value);
            } else {
                localStatePool[tmpTransitionList[i].account].lock  += tmpTransitionList[i].value;
                emit LocalStateNotify(tmpTransitionList[i].account, "lock-from-l2", tmpTransitionList[i].value);
            }

            lockTransitionPool[lockTransitionNum].chainID = tmpTransitionList[i].chainID;
            lockTransitionPool[lockTransitionNum].blockNum = tmpTransitionList[i].blockNum;
            lockTransitionPool[lockTransitionNum].account = tmpTransitionList[i].account;
            lockTransitionPool[lockTransitionNum].value = tmpTransitionList[i].value;
            lockTransitionPool[lockTransitionNum].style = tmpTransitionList[i].style;

            // calculate the unlock time
            lockTransitionPool[lockTransitionNum].unlockTime = block.number;

            lockTransitionNum ++;

        }

        return true;
    }

    //
    // verifyLocalState: use the local state record to verify the transition list, if the transition is illegal, return false
    function verifyLocalState(dt.L2Transition[] memory _transits) public view returns (bool) {
        return !_accumulatedDebitExceedsBalance(_transits);
    }

    // confirmLocalState confirm the locked state
    function confirmLocalStateWithBlockNum(uint256 _blockNum) public onlyRollup returns(uint256) {
        if (lockTransitionNum == 0) return 0;

        uint256 currentEnd = lockTransitionNum;

        for (uint256 i = currentEnd-1; i >=0 ; i --) {
            if (lockTransitionPool[i].blockNum == _blockNum) {
                // confirm
                if (lockTransitionPool[i].style == 1) {
                    localStatePool[lockTransitionPool[i].account].lock  -= lockTransitionPool[i].value;
                    emit LocalStateNotify(lockTransitionPool[i].account, "unlock-to-l2", lockTransitionPool[i].value);
                } else {
                    localStatePool[lockTransitionPool[i].account].value += lockTransitionPool[i].value;
                    localStatePool[lockTransitionPool[i].account].lock  -= lockTransitionPool[i].value;
                    emit LocalStateNotify(lockTransitionPool[i].account, "unlock-to-l1", lockTransitionPool[i].value);
                }

                // remove from pool
                lockTransitionPool[i].chainID = lockTransitionPool[currentEnd-1].chainID;
                lockTransitionPool[i].blockNum = lockTransitionPool[currentEnd-1].blockNum;
                lockTransitionPool[i].account = lockTransitionPool[currentEnd-1].account;
                lockTransitionPool[i].value = lockTransitionPool[currentEnd-1].value;
                lockTransitionPool[i].style = lockTransitionPool[currentEnd-1].style;
                currentEnd--;
            }

            if (i == 0) {
                break;
            }
        }

        lockTransitionNum = currentEnd;

        return 0;
    }

    // removeLocalState confirm the locked state
    function removeLocalStateWithBlockNum(uint256 _blockNum) public onlyRollup returns(uint256) {
        if (lockTransitionNum == 0) return 0;

        uint256 currentEnd = lockTransitionNum;

        for (uint256 i = currentEnd-1; i >=0 ; i --) {
            if (lockTransitionPool[i].blockNum >= _blockNum) {
                if (lockTransitionPool[i].style == 1) {
                    localStatePool[lockTransitionPool[i].account].value += lockTransitionPool[i].value;
                    localStatePool[lockTransitionPool[i].account].lock  -= lockTransitionPool[i].value;
                    emit LocalStateNotify(lockTransitionPool[i].account, "rollback-to-l1", lockTransitionPool[i].value);
                } else {
                    localStatePool[lockTransitionPool[i].account].lock  -= lockTransitionPool[i].value;
                    emit LocalStateNotify(lockTransitionPool[i].account, "rollback-to-l2", lockTransitionPool[i].value);
                }

                // remove from pool
                lockTransitionPool[i].chainID = lockTransitionPool[currentEnd-1].chainID;
                lockTransitionPool[i].blockNum = lockTransitionPool[currentEnd-1].blockNum;
                lockTransitionPool[i].account = lockTransitionPool[currentEnd-1].account;
                lockTransitionPool[i].value = lockTransitionPool[currentEnd-1].value;
                lockTransitionPool[i].style = lockTransitionPool[currentEnd-1].style;
                currentEnd--;
            }

            if (i == 0) {
                break;
            }
        }

        lockTransitionNum = currentEnd;

        return 0;
    }

    function checkRoot(dt.LocalTransition[] memory _transitions, bytes32 _root) public view returns(bool, bytes32) {
        bytes[] memory bytesTransitions = new bytes[](_transitions.length);

        for (uint i = 0; i < _transitions.length; i ++) {
            bytesTransitions[i] = dataTypes.LocalTransitionToBytes(_transitions[i]);
        }

        bytes32 tmpRoot = merkleUtils.getMerkleRoot(bytesTransitions);

        return ((tmpRoot == _root), tmpRoot);
    }

}
