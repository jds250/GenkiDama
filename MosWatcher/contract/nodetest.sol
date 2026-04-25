// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;
pragma experimental ABIEncoderV2;

import "solidity-rlp/contracts/RLPReader.sol";

/// @title MMR 验证器
contract MMRVerifier {
    // contract field
    address public superAddress;
    uint64 public MAX_UINT;

    //
    struct Block {
        string  fromAddr;
        uint256 diff;
        bytes   MRoot;
        uint256 Number;
        uint256 Nonce;
    }

    struct MMR_config {
        uint256 leafNum;
    }

    struct ProofRes {
        bytes32 h; // 哈希值
        uint256 td; // index
    }

    struct ProofElement {
        uint256 Cat;
        bool Right;
        uint64 LeafNum;
        ProofRes Res;
    }

    struct ProofInfo {
        bytes32 RootHash;
        uint256 RootDifficulty;
        uint64 LeafNumber;
        ProofElement[] Elems;
        uint64[] Checked;
    }

    struct ProofBlock {
        uint64 Number;
        uint256 AggrWeight;
    }

    struct VerifyElem {
        ProofRes Res;
        uint64 Index;
        uint64 LeafNumber;
    }


    constructor () {
        superAddress = msg.sender;
        MAX_UINT = 9999999;
    }

    // TODO: verify the proof-block with proof info
    function VerifyMMRProof(MMR_config memory _mmr, ProofInfo memory _info, ProofBlock[] memory _blocks) public returns (bool) {
        // 排序，去重，同时倒置。默认由链下执行器在提交前完成
        ProofElement memory proofElem;
        ProofElement memory right_node;
        VerifyElem[] memory nodes = new VerifyElem[](2*_info.Elems.length);
        bytes32 hash;

        // 设置索引
//        VerifyElem[] memory proofs = _info.Elems;
        uint front = 0;
        uint end = _info.Elems.length;
        uint blocksEnd = _blocks.length;
        uint nodesEnd = 0;
        // blockIndex = _blocks.length - 1;

        // 如果元素不存在或者直接是根，那么默认错误
        ProofElement memory root_elem = _info.Elems[end-1];
        end --;
        if (_info.Elems.length == 0 || root_elem.Cat != 0) {
            return false;
        }

        // 如果只有两个节点
        if (_info.Elems.length == 2) {
            if (_info.Elems[0].Cat == 2) {
                return _info.Elems[0].Res.h == _info.Elems[1].Res.h;
            }
            return false;
        }

        // 新建验证元素栈
        while (end > front) {
            // 如果还有元素
            // 从proofs(_info.Elems)最前方取一个元素
            ProofElement memory proof_elem = _info.Elems[front];
            front ++;

            if (proof_elem.Cat == 2) { // 待证明元素是叶子节点
                ProofBlock memory proof_block = _blocks[blocksEnd-1];
                blocksEnd --;
                uint64 number = proof_block.Number;

                if (nodesEnd > 0) {
                    (proofElem.Res.h, proofElem.Res.td) = getRoot(nodes);
                    uint256 left = proofElem.Res.td;
                    uint256 middle = root_elem.Res.td * proof_block.AggrWeight;
                    uint256 right = proofElem.Res.td + proof_elem.Res.td;

                    if (left > middle || right <= middle) {
                        return false;
                    }
                }

                if ((number%2 == 0) && (number != (root_elem.LeafNum-1))) {
                    right_node = _info.Elems[front];
                    front++;

                    if (right_node.Cat == 2 || right_node.Cat == 1) {
                        if (right_node.Cat == 2) {
                            blocksEnd --;
                        }
                    } else {
                        return false;
                    }

                    hash = merge2(proof_elem.Res.h, right_node.Res.h);
                    nodes[nodesEnd].Res.h = hash;
                    nodes[nodesEnd].Res.td = proof_elem.Res.td + right_node.Res.td;
                    nodes[nodesEnd].Index = number/2;
                    nodes[nodesEnd].LeafNumber = root_elem.LeafNum/2;

                    nodesEnd ++;
//                    nodes.push(tmpVerifyElem);

                } else {
//                    left_node = nodes[nodesEnd];
                    nodesEnd --;

                    hash = merge2(nodes[nodesEnd].Res.h, proof_elem.Res.h);
                    nodes[nodesEnd].Res.h = hash;
                    nodes[nodesEnd].Res.td = proof_elem.Res.td + nodes[nodesEnd].Res.td;
                    nodes[nodesEnd].Index = number/2;
                    nodes[nodesEnd].LeafNumber = root_elem.LeafNum/2;
                    nodesEnd ++;
//                    nodes.push(tmpVerifyElem);
                }

            } else if (proof_elem.Cat == 1) {
                if (proof_elem.Right) {
//                    left_node = nodes[nodesEnd];
                    nodesEnd --;

                    hash = merge2(nodes[nodesEnd].Res.h, proof_elem.Res.h);
                    nodes[nodesEnd].Res.h = hash;
                    nodes[nodesEnd].Res.td = proof_elem.Res.td + nodes[nodesEnd].Res.td;
                    nodes[nodesEnd].Index = nodes[nodesEnd].Index/2;
                    nodes[nodesEnd].LeafNumber = nodes[nodesEnd].LeafNumber/2;
                    nodesEnd++;
//                    nodes.push(tmpVerifyElem);
                } else {
                    nodes[nodesEnd].Res.h = proof_elem.Res.h;
                    nodes[nodesEnd].Res.td = proof_elem.Res.td;
                    nodes[nodesEnd].Index = 9999999;
                    nodes[nodesEnd].LeafNumber = 9999999;
                    nodesEnd ++;
                }
            } else if (proof_elem.Cat == 0){
                // do nothing
            } else {
                revert("invalid Cat...");
            }

            // 计算最终根


            while (nodesEnd > 1) {
                //  left_node
//                tmpVerifyElem2 = nodes.pop(); nodesEnd -1
//                tmpVerifyElem = nodes.pop(); nodesEnd - 2
                if (nodes[nodesEnd-1].Index == 9999999 || (nodes[nodesEnd-1].Index%2 != 1 && end > front)) {
//                    nodes.push(tmpVerifyElem);
//                    nodes.push(tmpVerifyElem2);
                    break;
                }

                hash = merge2(nodes[nodesEnd-2].Res.h, nodes[nodesEnd-1].Res.h);

//                nodes[nodesEnd].Res.h = hash;
//                nodes[nodesEnd].Res.td = tmpVerifyElem.Res.td + tmpVerifyElem2.Res.td;
//                nodes[nodesEnd].Index = tmpVerifyElem2.Index/2;
//                nodes[nodesEnd].LeafNumber = tmpVerifyElem2.LeafNumber/2;
//                nodesEnd ++;
                nodes[nodesEnd-2].Res.h = hash;
                nodes[nodesEnd-2].Res.td = nodes[nodesEnd-2].Res.td + nodes[nodesEnd-1].Res.td;
                nodes[nodesEnd-2].Index = nodes[nodesEnd-1].Index/2;
                nodes[nodesEnd-2].LeafNumber = nodes[nodesEnd-1].LeafNumber/2;
                nodesEnd --;
//                nodes.push(tmpVerifyElem);
            }
        }

        if (nodesEnd > 0) {
//            tmpVerifyElem =nodes.pop();
            return (hashCom(root_elem.Res.h, nodes[nodesEnd-1].Res.h) && (root_elem.Res.td == nodes[nodesEnd-1].Res.td));
        }

        return false;
    }

    // getRoot 获取验证元素的根
    function getRoot(VerifyElem[] memory nodes) public returns(bytes32, uint256) {
        // VerifyElem[] memory tmp;
        uint end = nodes.length;
        VerifyElem memory node1;
        VerifyElem memory node2;

        while (end > 1) {
            node1 = nodes[end-1];
            end --;
            node2 = nodes[end-1];
            end --;
            bytes32 hash = merge2(node1.Res.h, node2.Res.h);

            // 记录新根
            nodes[end].Index = 9999999;
            nodes[end].LeafNumber = 9999999;
            nodes[end].Res.h = hash;
            nodes[end].Res.td = node1.Res.td + node2.Res.td;
            end ++;
        }

        if (end>=1) {
            return (nodes[0].Res.h,nodes[0].Res.td);
        }
        return (bytes32(""),0);
    }

    // TODO: get the random block
    function RandomSample(uint256 _leafNumber, uint256 _cnt) public pure returns(uint256[] memory) {
        uint256[] memory res = new uint256[](_cnt);
        return res;
    }

    // hashCom
    function hashCom(bytes32 a, bytes32 b) public pure returns(bool) {
        return a == b;
    }

    // other 其他辅助函数
    function bytesCom(bytes memory a, bytes memory b) public pure returns(bool) {
        if(a.length == b.length){
            for(uint i = 0;i<a.length;++i){
                if(a[i] != b[i])
                    return false;
            }
        }else{
            return false;
        }
        return true;
    }

    function merge2(bytes32 h1, bytes32 h2) public returns(bytes32) {
        bytes memory array68 = new bytes(68);

        // 设置特殊字段
        array68[0] = 0xf8;
        array68[1] = 0x42;
        array68[2] = 0xa0;
        array68[35] = 0xa0;

        // 复制 h1 和 h2
        for (uint i = 0; i < 32; i ++) {
            array68[3 + i] = h1[i];
            array68[36 + i] = h2[i];
        }

        return keccak256(array68);
    }

}