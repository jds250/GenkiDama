// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;

/* Internal Imports */
import {DataTypes} from "./DataTypes.sol";
import {MerkleUtils} from "./MerkleUtils.sol";
import {MPTVerifier} from "./MPTVerifier.sol";

/// @title 多链执行合约
contract FullRollup {

    /*** Fields ***/
    /* Contract Instance */
    // The Merkle Tree library (currently a contract for ease of testing)
    MerkleUtils myMU;
    // Data Type
    DataTypes myDT;
    // MPT trie Verifier
    MPTVerifier myMPTV;

    /* Const */
    bytes32 public constant ZERO_BYTES32 = 0x0000000000000000000000000000000000000000000000000000000000000000;
    // TODO: Set a reasonable wait period
    uint256 public constant WITHDRAW_WAIT_PERIOD = 4;
    // dispute time
    uint256 public constant DISPUTE_TIME = 1;

    /* rollup system state */
    // All the blocks!
    uint256 public chainID;
    DataTypes.Block[] public L2Blocks;
    uint256 public confirmedNum;

    mapping(uint256=>DataTypes.ChallengeState) public challengePool;
    uint256 public challengeSize = 0;

    // authority
    address public committerAddress;
    address public superAddress;

    /* Events */
    // to notify the block state
    // suc: 0-success 1-fail
    // style: 1-commit 2-confirm 3-rollback 4-challenged
    // desc = style-suc-desc
    event RollupBlockNotify(uint256 blockNumber, string desc);
    event ErrorNotify(uint256 state, string error);
    event LocalStateNotify(string account, string style, uint256 value);

    /***************
     * Constructor *
     **************/
    constructor(
        uint256 _chainID,
        address _committerAddress,
        address _MerkleUtilsAddress,
        address _DataTypesAddress,
        address _MPTVerifierAddress
    ) {
        // config
        chainID = _chainID;
        committerAddress = _committerAddress;
        superAddress = msg.sender;

        // contract init
        myMU = MerkleUtils(_MerkleUtilsAddress);
        myDT = DataTypes(_DataTypesAddress);
        myMPTV = MPTVerifier(_MPTVerifierAddress);

        // state init
        confirmedNum = 0;

        // local state init
        lockTransitionNum = 0;
    }

    modifier onlySuper() {
        require(
            msg.sender == address(superAddress),
            "Only super admin may perform action"
        );
        _;
    }

    modifier onlyCommitter() {
        require(
            msg.sender == address(committerAddress),
            "Only committer may perform action"
        );
        _;
    }

    /* Methods */
    function pruneBlocksAfter(uint256 _blockNumber) internal {
        for (uint256 i = _blockNumber; i < L2Blocks.length; i++) {
            delete L2Blocks[i];
        }
    }

    function getCurrentBlockNumber() public view returns (uint256) {
        return L2Blocks.length - 1;
    }

    function setCommitterAddress(address _committerAddress)
    external
    {
        committerAddress = _committerAddress;
    }

    // get commit size
    function BlockLen() public view returns(uint256) {
        return L2Blocks.length;
    }

    // commitBlock: TxUpload and commit block verify
    /**
     * Commits a new block which is then rolled up.
     */
    function commitBlock(
        uint256 _blockNumber,
        bytes32 _txRoot,
        bytes32 _localTransitionRoot,
        DataTypes.L2Transaction[] calldata _transactions,
        DataTypes.L2Transition[] calldata _transitionList
    ) external returns (bool) {
        // emit ErrorNotify(1, "This is a test");
        if (_blockNumber != L2Blocks.length) {
            emit ErrorNotify(1, "Illegal block number");
            return false;
        }

        // verify transaction root
        bytes32 tmpRoot = myDT.TransactionsListHash(_transactions);
        if (tmpRoot != _txRoot) {
            emit ErrorNotify(1, "Illegal transaction root");
            return false;
        }

        // // verify local state root
        // tmpRoot = myDT.L2TransitionListHash(_transitionList);
        // if (tmpRoot != _localTransitionRoot) {
        //     emit ErrorNotify(1, "Illegal local state root");
        //     return false;
        // }

        DataTypes.Block memory rollupBlock;
        rollupBlock.height = _blockNumber;
        rollupBlock.blockSize = _transactions.length;
        rollupBlock.txRoot = _txRoot;
        rollupBlock.localTransitionRoot = _localTransitionRoot;
        rollupBlock.commitTime = block.number;
        L2Blocks.push(rollupBlock);

        // lock local state
        bool res = verifyAndLockLocalState(_transitionList, _blockNumber, block.number);

        if (!res) {
            emit ErrorNotify(1, "Illegal state verify");
            return false;
        }

        emit RollupBlockNotify(_blockNumber, "1-0-commit success");
        return true;
    }

    // checkCommit: confirm the block
    function checkCommit() public {
        // check if the commit is enough to confirm
        // if so, confirm the
        uint256 start = block.number-DISPUTE_TIME;
        for (uint256 i = confirmedNum; i < L2Blocks.length; i ++) {
            if (L2Blocks[i].commitTime <= start) {
                uint256 res = confirmLocalStateWithBlockNum(L2Blocks[i].height);

                if (res == 0) {
                    emit RollupBlockNotify(uint256(i), "2-0-commit success");
                } else {
                    emit RollupBlockNotify(uint256(i), "2-1-commit fail");
                    revert();
                }

            } else {
                confirmedNum = i;
                break;
            }
        }
    }

    // challengeBlockExecution: Challenge a block and begin the Verification of transactions
    function ExecutionChallenge(
        uint256 _targetBlockHeight,
        DataTypes.L2Transaction[] memory _transactions
    ) public {
        // require( _targetBlockHeight > L2Blocks.length , "Block not Exist");
        // require( _targetBlockHeight <= confirmedNum , "Block is Confirmed");
        // DataTypes.Block memory targetBlock = L2Blocks[_targetBlockHeight];

        if (_targetBlockHeight >= L2Blocks.length) {
            emit ErrorNotify(2, "don exist");
            return ;
        }

        bytes32 tmpRoot = myDT.TransactionsListHash(_transactions);
        if (tmpRoot != L2Blocks[_targetBlockHeight].txRoot) {
            emit ErrorNotify(1, "Illegal transaction root");
            return;
        }

        bool right;
        string memory desc;
        (right, desc) = VerifyTransactions(_transactions, L2Blocks[_targetBlockHeight].localTransitionRoot);

        if (!right) {
            emit RollupBlockNotify(_targetBlockHeight, strConcat("4-0-", desc));
            rollBackBlock(_targetBlockHeight);
        }
        emit RollupBlockNotify(_targetBlockHeight, "4-1-The Challenged Block is legal");
    }

    // TODO: rollupBackBlock roll back commitment _targetBlockHeight and the following blocks
    // TODO: unrealize
    function rollBackBlock(uint256 _targetBlockHeight) public {

        uint256 res = removeLocalStateWithBlockNum(_targetBlockHeight);
        emit RollupBlockNotify(_targetBlockHeight, "remove success");
    }

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



    /*
       Local state management
    */
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

    // depositCoin deposit and generate a localStateRecord
    function depositCoin(string memory _account, uint256 _value) public {
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
    function refundCoin(string memory _account, uint256 _value) public {
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
    function verifyAndLockLocalState(dt.L2Transition[] memory _transits, uint256 _blockNum, uint256 _lockTime) public returns (bool) {
        dt.LocalTransition[] memory tmpTransitionList = new dt.LocalTransition[](
            _transits.length
        );
        uint256 tmpNum = 0;

        // decode all LocalState and verify
        for (uint256 i = 0; i < _transits.length; i++) {

            if (_transits[i].style == 1
                && _transits[i].value > localStatePool[_transits[i].account].value) {
                emit ErrorNotify(1, "illegal transition: account balance is not enough");
                return false;
            }

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
        for (uint256 i = 0; i < _transits.length; i++) {

            if (_transits[i].chainID == chainID && _transits[i].style == 1
                && _transits[i].value > localStatePool[_transits[i].account].value) {
                return false;
            }

        }

        return true;
    }

    // confirmLocalState confirm the locked state
    function confirmLocalStateWithBlockNum(uint256 _blockNum) public returns(uint256) {
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
    function removeLocalStateWithBlockNum(uint256 _blockNum) public returns(uint256) {
        if (lockTransitionNum == 0) return 0;

        uint256 currentEnd = lockTransitionNum;

        for (uint256 i = currentEnd-1; i >=0 ; i --) {
            if (lockTransitionPool[i].blockNum >= _blockNum) {
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


    /*
       Transaction Executor
    */
    function VerifyTransactions(
        DataTypes.L2Transaction[] memory _transactions,
        bytes32 _tsRoot
    ) public view returns (bool, string memory) {
        // get L2Transaction
        DataTypes.L2Transition[] memory l2Transitions = GetL2Transitions(_transactions);

        // verify local Transition with root
        bytes32 tmpRoot = myDT.L2TransitionListHash(l2Transitions);
        if (_tsRoot != tmpRoot) {
            return (false, "targetBlock.localStateRoot is Wrong");
        }

        // verify the local transition with local state
        bool result = verifyLocalState(l2Transitions);

        if (result) {
            return (true, "Proof Success!");
        } else {
            return (false, "no enough token");
        }
    }


    // GetL2Transitions Translate Transaction to L2 Transition
    function GetL2Transitions(DataTypes.L2Transaction[] memory _transactions) public pure returns (DataTypes.L2Transition[] memory) {
        uint256 cnt = 2 * _transactions.length;

        DataTypes.L2Transition[] memory resTransit = new DataTypes.L2Transition[](cnt);

        uint256 index = 0;
        for (uint256 i = 0; i < _transactions.length; i ++) {
            // 添加输入状态转化
            if (_transactions[i].style == 1) {
                resTransit[index].chainID = 1;
                resTransit[index+1].chainID = 2;
            } else {
                resTransit[index].chainID = 2;
                resTransit[index+1].chainID = 1;
            }

            // 添加输入状态转化
            resTransit[index].account = _transactions[i].fromAddr;
            resTransit[index].value = _transactions[i].value;
            resTransit[index].style = 1;

            // 添加输出状态转化
            resTransit[index].account = _transactions[i].toAddr;
            resTransit[index].value = _transactions[i].value;
            resTransit[index].style = 2;

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
}