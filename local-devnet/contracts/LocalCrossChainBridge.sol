// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract LocalCrossChainBridge {
    uint256 public immutable chainId;
    uint256 public nextMessageId;

    mapping(address => uint256) public balances;
    mapping(bytes32 => bool) public processedMessages;

    event Deposited(address indexed account, uint256 amount);
    event OutboundCommitted(
        bytes32 indexed messageId,
        uint256 indexed toChainId,
        address indexed from,
        address to,
        uint256 amount
    );
    event InboundApplied(
        bytes32 indexed messageId,
        uint256 indexed fromChainId,
        address indexed to,
        uint256 amount
    );

    constructor(uint256 _chainId) {
        chainId = _chainId;
    }

    function deposit(address account, uint256 amount) external {
        require(account != address(0), "account is zero");
        require(amount > 0, "amount is zero");

        balances[account] += amount;
        emit Deposited(account, amount);
    }

    function commitOutbound(
        uint256 toChainId,
        address to,
        uint256 amount
    ) external returns (bytes32 messageId) {
        require(to != address(0), "to is zero");
        require(amount > 0, "amount is zero");
        require(balances[msg.sender] >= amount, "insufficient balance");

        balances[msg.sender] -= amount;

        messageId = keccak256(
            abi.encodePacked(chainId, toChainId, msg.sender, to, amount, nextMessageId)
        );
        nextMessageId += 1;

        emit OutboundCommitted(messageId, toChainId, msg.sender, to, amount);
    }

    function applyInbound(
        bytes32 messageId,
        uint256 fromChainId,
        address from,
        address to,
        uint256 amount
    ) external {
        require(!processedMessages[messageId], "message already processed");
        require(to != address(0), "to is zero");
        require(amount > 0, "amount is zero");

        processedMessages[messageId] = true;
        balances[to] += amount;

        emit InboundApplied(messageId, fromChainId, to, amount);
    }
}
