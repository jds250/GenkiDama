// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;

/* Internal Imports */
import {DataTypes} from "./DataTypes.sol";
import {MerkleUtils} from "./MerkleUtils.sol";
import {TransactionExecutor} from "./TransactionExecutor.sol";
import {MPTVerifier} from "./MPTVerifier.sol";
import {LocalStateManager} from "./LocalStateManager.sol";

/// @title 多链执行合约
contract CrossRollup {

    /*** Fields ***/
    /* Contract Instance */
    // The state transition evaluator
    TransactionExecutor myTE;
    // The Merkle Tree library (currently a contract for ease of testing)
    MerkleUtils myMU;
    // Data Type
    DataTypes myDT;
    // MPT trie Verifier
    MPTVerifier myMPTV;
    // Local State Manager
    LocalStateManager myLSM;

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

    /***************
     * Constructor *
     **************/
    constructor(
        uint256 _chainID,
        address _committerAddress,
        address _TransactionExecutorAddress,
        address _MerkleUtilsAddress,
        address _DataTypesAddress,
        address _MPTVerifierAddress,
        address _LocalStateManagerAddress
    ) {
        // config
        chainID = _chainID;
        committerAddress = _committerAddress;
        superAddress = msg.sender;

        // contract init
        myTE = TransactionExecutor(_TransactionExecutorAddress);
        myMU = MerkleUtils(_MerkleUtilsAddress);
        myDT = DataTypes(_DataTypesAddress);
        myMPTV = MPTVerifier(_MPTVerifierAddress);
        myLSM = LocalStateManager(_LocalStateManagerAddress);

        // state init
        confirmedNum = 0;
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
        // every commit will check the old committed block
        // checkCommit();

        // require(
        //     msg.sender == superAddress,
        //     "Only the superAdmin may submit blocks"
        // );
        // require(_blockNumber == L2Blocks.length, "Illegal block number");


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
        bool success;
        bytes memory returnData;
        (success, returnData) = address(myLSM).call(
            abi.encodeWithSelector(
                myLSM
                .verifyAndLockLocalState
                .selector,
                _transitionList, _blockNumber, block.number
            )
        );

        bool res;
        if (success) {
            (res) = abi.decode(
                returnData,
                (bool)
            );
        }

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
                // confirm the local state
                bool success;
                bytes memory returnData;
                (success, returnData) = address(myLSM).call(
                    abi.encodeWithSelector(
                        myLSM
                        .confirmLocalStateWithBlockNum
                        .selector,
                        L2Blocks[i].height
                    )
                );

                if (success) {
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
        (right, desc) = myTE.VerifyTransactions(_transactions, L2Blocks[_targetBlockHeight].localTransitionRoot);

        if (!right) {
            emit RollupBlockNotify(_targetBlockHeight, strConcat("4-0-", desc));
            rollBackBlock(_targetBlockHeight);
        }
        emit RollupBlockNotify(_targetBlockHeight, "4-1-The Challenged Block is legal");
    }

    // TODO: the realization of interactive communication
    function challengeArgu() public {

    }

    // TODO: Verifiy cross chain inconsistent
    function verifyCrossBlock() public {

    }


    // TODO: rollupBackBlock roll back commitment _targetBlockHeight and the following blocks
    // TODO: unrealize
    function rollBackBlock(uint256 _targetBlockHeight) public {
        // build local state
        bool success;
        bytes memory returnData;
        (success, returnData) = address(myLSM).call(
            abi.encodeWithSelector(
                myLSM
                .removeLocalStateWithBlockNum
                .selector,
                _targetBlockHeight
            )
        );
        if (! success) {
            emit RollupBlockNotify(_targetBlockHeight, "3-1-The Challenged Block is legal");
        } else {
            emit RollupBlockNotify(_targetBlockHeight, "3-0-The Challenged Block is legal");
        }

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


}