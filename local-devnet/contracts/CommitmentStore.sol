// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

contract CommitmentStore {
    bytes32 public commitmentRoot;

    function setCommitmentRoot(bytes32 newRoot) external {
        commitmentRoot = newRoot;
    }
}
