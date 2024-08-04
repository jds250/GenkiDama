// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package trie

import (
	"github.com/HyperService-Consortium/go-rlp"
	"github.com/syndtr/goleveldb/leveldb"
	"golang.org/x/crypto/sha3"
	"hash"
	"sync"
)

type hasher struct {
	tmp    sliceBuffer
	sha    keccakState
	onleaf LeafCallback
}

// keccakState wraps sha3.state. In addition to the usual hash methods, it also supports
// Read to get a variable amount of data from the hash state. Read is faster than Sum
// because it doesn't copy the internal state, but also modifies the internal state.
type keccakState interface {
	hash.Hash
	Read([]byte) (int, error)
}

type sliceBuffer []byte

func (b *sliceBuffer) Write(data []byte) (n int, err error) {
	*b = append(*b, data...)
	return len(data), nil
}

func (b *sliceBuffer) Reset() {
	*b = (*b)[:0]
}

// hashers live in a global db.
var hasherPool = sync.Pool{
	New: func() interface{} {
		return &hasher{
			tmp: make(sliceBuffer, 0, 550), // cap is as large as a full FullNode.
			sha: sha3.NewLegacyKeccak256().(keccakState),
		}
	},
}

func newHasher(onleaf LeafCallback) *hasher {
	h := hasherPool.Get().(*hasher)
	h.onleaf = onleaf
	return h
}

func returnHasherToPool(h *hasher) {
	hasherPool.Put(h)
}

// hash collapses a node down into a hash node, also returning a copy of the
// original node initialized with the computed hash to replace the original one.
func (h *hasher) hash(n node, db *leveldb.DB, force bool) (node, node, error) {
	// If we're not storing the node, just hashing, use available cached data
	//计算节点的哈希值，首先检查节点的哈希值是否已经缓存，如果是，则返回缓存的哈希值。否则，方法调用hashChildren来计算节点子节点的哈希值。如果计算成功，方法调用store将节点存储在leveldb数据库中，
	//然后缓存节点的哈希值以供以后使用。如果force参数为真，则节点总是存储在数据库中，即使其哈希值已经被缓存。如果为假，则只有在节点的哈希值未被缓存时才将其存储在数据库中。
	//该方法返回两个node类型：哈希节点和原始节点。哈希节点是节点的哈希值存储在leveldb数据库中的节点。原始节点是输入节点，它可能已经缓存了其哈希值。
	//如果在计算过程中发生错误，该方法返回一个错误。否则，它返回哈希节点、缓存节点（如果有的话）和一个空错误
	if hash, dirty := n.cache(); hash != nil {
		if db == nil {
			//fmt.Println("db")
			return hash, n, nil
		}
		if !dirty {
			switch n.(type) {
			case *FullNode, *ShortNode:
				//fmt.Println("case1:", hash)
				return hash, hash, nil
			default:
				//fmt.Println("case2")
				return hash, n, nil
			}
		}
	}
	// Trie not processed yet or needs storage, walk the children
	collapsed, cached, err := h.hashChildren(n, db)
	if err != nil {
		return HashNode{}, n, err
	}
	hashed, err := h.store(collapsed, db, force)
	//fmt.Println("stored:", hashed, "collapsed:", collapsed)
	if err != nil {
		return HashNode{}, n, err
	}
	// Cache the hash of the node for later reuse and remove
	// the dirty flag in commit mode. It's fine to assign these values directly
	// without copying the node first because hashChildren copies it.
	cachedHash, _ := hashed.(HashNode)
	switch cn := cached.(type) {
	case *ShortNode:
		cn.flags.hash = cachedHash
		if db != nil {
			cn.flags.dirty = false
		}
	case *FullNode:
		cn.flags.hash = cachedHash
		if db != nil {
			cn.flags.dirty = false
		}
	}
	return hashed, cached, nil
}

// hashChildren replaces the children of a node with their hashes if the encoded
// size of the child is larger than a hash, returning the collapsed node as well
// as a replacement for the original node with the child hashes cached in.
func (h *hasher) hashChildren(original node, db *leveldb.DB) (node, node, error) {
	var err error

	switch n := original.(type) {
	case *ShortNode:
		// Hash the short node's child, caching the newly hashed subtree
		collapsed, cached := n.copy(), n.copy()
		collapsed.Key = hexToCompact(n.Key)
		cached.Key = CopyBytes(n.Key)

		if _, ok := n.Val.(ValueNode); !ok {
			collapsed.Val, cached.Val, err = h.hash(n.Val, db, false)
			if err != nil {
				return original, original, err
			}
		}
		return collapsed, cached, nil

	case *FullNode:
		// Hash the full node's children, caching the newly hashed subtrees
		collapsed, cached := n.copy(), n.copy()

		for i := 0; i < 16; i++ {
			if n.Children[i] != nil {
				collapsed.Children[i], cached.Children[i], err = h.hash(n.Children[i], db, false)
				if err != nil {
					return original, original, err
				}
			}
		}
		cached.Children[16] = n.Children[16]
		return collapsed, cached, nil

	default:
		// Value and hash nodes don't have children so they're left as were
		return n, original, nil
	}
}

// store hashes the node n and if we have a storage layer specified, it writes
// the key/value pair to it and tracks any node->child references as well as any
// node->external trie references.
func (h *hasher) store(n node, db *leveldb.DB, force bool) (node, error) {
	// Don't store hashes or empty nodes.
	if _, isHash := n.(HashNode); n == nil || isHash {
		return n, nil
	}
	// Generate the RLP encoding of the node
	h.tmp.Reset()
	if err := rlp.Encode(&h.tmp, n); err != nil {
		panic("encode error: " + err.Error())
	}
	if len(h.tmp) < 32 && !force {
		return n, nil // Nodes smaller than 32 bytes are stored inside their parent
	}
	// Larger nodes are replaced by their hash and stored in the database.
	hashnode, _ := n.cache()
	if hashnode == nil {
		hashnode = h.makeHashNode(h.tmp)
	}

	if db != nil {
		// We are pooling the trie nodes into an intermediate memory cache
		//var hnode = hashnode
		hashnode := BytesToHash(hashnode)

		var rwLock sync.RWMutex
		rwLock.Lock()
		db.Put(hashnode.Bytes(), h.tmp, nil)
		//fmt.Println("stored:", hnode)
		rwLock.Unlock()

		// Track external references from account->storage trie
		if h.onleaf != nil {
			switch n := n.(type) {
			case *ShortNode:
				if child, ok := n.Val.(ValueNode); ok {
					h.onleaf(child, hashnode)
				}
			case *FullNode:
				for i := 0; i < 16; i++ {
					if child, ok := n.Children[i].(ValueNode); ok {
						h.onleaf(child, hashnode)
					}
				}
			}
		}
	}
	return hashnode, nil
}

func (h *hasher) makeHashNode(data []byte) HashNode {
	n := make(HashNode, h.sha.Size())
	h.sha.Reset()
	h.sha.Write(data)
	h.sha.Read(n)
	return n
}
