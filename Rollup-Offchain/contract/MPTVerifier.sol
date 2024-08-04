// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;

import {RLPReader} from "solidity-rlp/contracts/RLPReader.sol";

/// @title Merkle Patricia Trie Proof Verifier.
/// @author Noah Figueras.
/// @notice You can use this contract to verify any MPT Proof.
/// @dev Used specifically to verify proofs from @ethereumjs/trie.

///  Documentation:
///  - https://github.com/pipeos-one/goldengate/blob/master/contracts/contracts/lib/MPT.sol
///  - https://easythereentropy.wordpress.com/2014/06/04/understanding-the-ethereum-trie/
///  - https://ouvrard-pierre-alain.medium.com/merkle-proof-verification-for-ethereum-patricia-tree-48f29658eec

contract MPTVerifier {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;

    /// @param expectedRoot MPT root.
    /// @param key keccak256 encoded key of value.
    /// @param proof Proof of key/value pair from MPT.
    /// @param keyIndex used for recursion on the key.
    /// @param proofIndex used for recursion on the proof.
    /// @param expectedValue RLP encoded value.
    struct MerkleProof {
        bytes32 expectedRoot;
        bytes key;
        bytes[] proof;
        uint256 keyIndex;
        uint256 proofIndex;
        bytes expectedValue;
    }

    /// @param data The paramaters needed for MPT verification.
    /// @return true if the proof is correct.
    function verifyProof(MerkleProof memory data) public view returns (bool) {
        bytes memory node = data.proof[data.proofIndex];
        RLPReader.RLPItem[] memory dec = node.toRlpItem().toList();
        require(keccak256(node) == data.expectedRoot, "invalid root hash");

        uint numItems = dec.length;

        if(numItems == 17) {
            // Branch Node
            if(data.keyIndex >= data.key.length) {
                if(keccak256(dec[16].toBytes()) == keccak256(data.expectedValue)) {
                    return true;
                }
            } else {
                uint8 index = uint8(bufferToNibbles(data.key)[data.keyIndex]);
                bytes32 newExpectedRoot = bytes32(dec[index].toBytes());
                if(newExpectedRoot.length != 0) {
                    data.expectedRoot = newExpectedRoot;
                    data.keyIndex += 1;
                    data.proofIndex += 1;
                    return verifyProof(data);
                }
            }
        } else if(numItems == 2) {
            bytes memory nodeKey = dec[0].toBytes();
            bytes memory nodeValue = dec[1].toBytes();
            bytes1[] memory restKey = bufferToNibbles(nodeKey);
            bytes1[] memory actualKey = bufferToNibbles(data.key);
            uint8 prefix = uint8(restKey[0]);

            if(prefix == 2) {
                // Even Leaf Node
                if(keccak256(slice(actualKey, data.keyIndex)) == keccak256(slice(restKey,2)) &&
                    keccak256(data.expectedValue) == keccak256(nodeValue)
                ) {
                    return true;
                }
            } else if(prefix == 3) {
                // Odd Leaf Node
                if(keccak256(slice(actualKey, data.keyIndex)) == keccak256(slice(restKey,1)) &&
                    keccak256(data.expectedValue) == keccak256(nodeValue)
                ) {
                    return true;
                }
            } else if(prefix == 0) {
                // Even Extension Node
                bytes memory sharedNibbles = slice(restKey, 2);
                uint extensionLength = sharedNibbles.length;
                if(keccak256(sharedNibbles) == keccak256(slice(actualKey, data.keyIndex, data.keyIndex + extensionLength))) {
                    data.expectedRoot = bytes32(dec[1].toBytes());
                    data.keyIndex += extensionLength;
                    data.proofIndex += 1;
                    return verifyProof(data);
                }
            } else if(prefix == 1) {
                // Odd Extension Node
                bytes memory sharedNibbles = slice(restKey, 1);
                uint extensionLength = sharedNibbles.length;
                if(keccak256(sharedNibbles) == keccak256(slice(actualKey, data.keyIndex, data.keyIndex + extensionLength))) {
                    data.expectedRoot = bytes32(dec[1].toBytes());
                    data.keyIndex += extensionLength;
                    data.proofIndex += 1;
                    return verifyProof(data);
                }
            } else {
                // This shouldn't be reached if the proof has the correct format
                revert("Invalid proof");
            }
        }
        return data.expectedValue.length == 0 ? true : false;
    }

    /// @dev Conversion of node from bytes to nibbles
    /// @param _key node in bytes
    /// @return nibbles array of bytes
    function bufferToNibbles(bytes memory _key) private pure returns(bytes1[] memory nibbles) {
        nibbles = new bytes1[](_key.length * 2);
        for(uint i = 0; i < _key.length; i++) {
            uint q = i * 2;
            nibbles[q] = _key[i] >> 4;
            ++q;
            nibbles[q] = bytes1(uint8(_key[i]) % 16);
        }
    }

    /// @dev Slices array of nibbles
    /// @param _key nibbles
    /// @param start starting point in the array
    /// @return result concatenated bytes of nibbles
    function slice(bytes1[] memory _key, uint start) private pure returns(bytes memory result) {
        for(uint i = start; i < _key.length; i++) {
            result = abi.encodePacked(result, _key[i]);
        }
    }

    /// @dev Slices array of nibbles
    /// @param _key nibbles
    /// @param start starting point in the array
    /// @param end ending point in the array
    /// @return result concatenated bytes of nibbles
    function slice(bytes1[] memory _key, uint start, uint end) private pure returns(bytes memory result) {
        for(uint i = start; i < end; i++) {
            result = abi.encodePacked(result, _key[i]);
        }
    }
}
