// Copyright 2014 The go-ethereum Authors
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

// Package trie implements Merkle Patricia Tries.
package trie

import (
	"bytes"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

var Count1, Count2 int = 0, 0

var (
	// emptyRoot is the known root hash of an empty trie.
	emptyRoot = HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")

	// emptyState is the known hash of an empty state trie entry.
	emptyState = Keccak256Hash(nil)
)

// LeafCallback is a callback type invoked when a trie operation reaches a leaf
// node. It's used by state sync and commit to allow handling external references
// between account and storage tries.
type LeafCallback func(leaf []byte, parent Hash) error

// Trie is a Merkle Patricia Trie.
// The zero value is an empty trie with no database.
// Use New to create a trie that sits on top of a database.
//
// Trie is not safe for concurrent use.
type Trie struct {
	db   *leveldb.DB
	root node
}

// newFlag returns the cache flag value for a newly created node.
func (t *Trie) newFlag() nodeFlag {
	return nodeFlag{dirty: true}
}

// New creates a trie with an existing root node from db.
//
// If root is the zero hash or the sha3 hash of an empty string, the
// trie is initially empty and does not require a database. Otherwise,
// New will panic if db is nil and returns a MissingNodeError if root does
// not exist in the database. Accessing the trie loads nodes from db on demand.
func NewTrie(root Hash, db *leveldb.DB) (*Trie, error) {
	if db == nil {
		panic("trie.New called without a database")
	}
	trie := &Trie{
		db: db,
	}
	if root != (Hash{}) && root != emptyRoot {
		rootnode, err := trie.resolveHash(root[:], nil)
		if err != nil {
			fmt.Println("NewTrie err")
			return nil, err
		}
		trie.root = rootnode
	}
	return trie, nil
}

// // NodeIterator returns an iterator that returns nodes of the trie. Iteration starts at
// // the key after the given start key.
// func (t *Trie) NodeIterator(start []byte) NodeIterator {
// 	return newNodeIterator(t, start)
// }

// Get returns the value for key stored in the trie.
// The value bytes must not be modified by the caller.
func (t *Trie) Get(key []byte) []byte {
	res, err := t.TryGet(key)
	if err != nil {
		panic(fmt.Sprintf("Unhandled trie error: %v", err))
	}
	return res
}

// TryGet returns the value for key stored in the trie.
// The value bytes must not be modified by the caller.
// If a node was not found in the database, a MissingNodeError is returned.
func (t *Trie) TryGet(key []byte) ([]byte, error) {
	key = keybytesToHex(key)
	value, newroot, didResolve, err := t.tryGet(t.root, key, 0)
	if err == nil && didResolve {
		t.root = newroot
	}
	return value, err
}

func (t *Trie) tryGet(origNode node, key []byte, pos int) (value []byte, newnode node, didResolve bool, err error) {
	switch n := (origNode).(type) {
	case nil:
		return nil, nil, false, nil
	case ValueNode:
		return n, n, false, nil
	case *ShortNode:
		if len(key)-pos < len(n.Key) || !bytes.Equal(n.Key, key[pos:pos+len(n.Key)]) {
			// key not found in trie
			return nil, n, false, nil
		}
		value, newnode, didResolve, err = t.tryGet(n.Val, key, pos+len(n.Key))
		if err == nil && didResolve {
			n = n.copy()
			n.Val = newnode
		}
		return value, n, didResolve, err
	case *FullNode:
		value, newnode, didResolve, err = t.tryGet(n.Children[key[pos]], key, pos+1)
		if err == nil && didResolve {
			n = n.copy()
			n.Children[key[pos]] = newnode
		}
		return value, n, didResolve, err
	case HashNode:
		child, err := t.resolveHash(n, key[:pos])
		if err != nil {
			return nil, n, true, err
		}
		value, newnode, _, err := t.tryGet(child, key, pos)
		return value, newnode, true, err
	default:
		panic(fmt.Sprintf("%T: invalid node: %v", origNode, origNode))
	}
}

func (t *Trie) Prove(key []byte) [][]byte {
	res, err := t.TryProve(key)
	if err != nil {
		panic(fmt.Sprintf("Unhandled trie error: %v", err))
	}
	return res
}

// TryGet returns the value for key stored in the trie.
// The value bytes must not be modified by the caller.
// If a node was not found in the database, a MissingNodeError is returned.
func (t *Trie) TryProve(key []byte) ([][]byte, error) {
	key = keybytesToHex(key)
	trieHash, err := t.Commit(nil)
	var pureTrie *Trie
	pureTrie, err = NewTrie(trieHash, t.db)
	if err != nil {
		return nil, err
	}
	var pbt []byte
	pbt, err = t.db.Get(trieHash[:], nil)
	if err != nil {
		return nil, err
	}
	var tailProof [][]byte
	tailProof, _, _, err = pureTrie.tryProve(pureTrie.root, key, 0)
	return append([][]byte{trieHash[:], pbt}, tailProof...), err
}

func (t *Trie) tryProve(origNode node, key []byte, pos int) (proof [][]byte, newnode node, didResolve bool, err error) {
	switch n := (origNode).(type) {
	case nil:
		return nil, nil, false, nil
	case ValueNode:
		return [][]byte{}, n, false, nil
	case *ShortNode:
		if len(key)-pos < len(n.Key) || !bytes.Equal(n.Key, key[pos:pos+len(n.Key)]) {
			// key not found in trie
			return [][]byte{}, n, false, nil
		}
		proof, newnode, didResolve, err = t.tryProve(n.Val, key, pos+len(n.Key))
		if err == nil && didResolve {
			n = n.copy()
			n.Val = newnode
		}
		return proof, n, didResolve, err
	case *FullNode:
		proof, newnode, didResolve, err = t.tryProve(n.Children[key[pos]], key, pos+1)
		if err == nil && didResolve {
			n = n.copy()
			n.Children[key[pos]] = newnode
		}
		return proof, n, didResolve, err
	case HashNode:
		child, err := t.resolveHash(n, key[:pos])
		if err != nil {
			return nil, n, true, err
		}
		var pbt []byte
		pbt, err = t.db.Get(BytesToHash(n).Bytes(), nil)
		if err != nil {
			return nil, n, true, err
		}
		tailProof, newnode, _, err := t.tryProve(child, key, pos)
		return append([][]byte{n, pbt}, tailProof...), newnode, true, err
	default:
		panic(fmt.Sprintf("%T: invalid node: %v", origNode, origNode))
	}
}

// Update associates key with value in the trie. Subsequent calls to
// Get will return value. If value has length zero, any existing value
// is deleted from the trie and calls to Get will return nil.
//
// The value bytes must not be modified by the caller while they are
// stored in the trie.
func (t *Trie) Update(key, value []byte) {
	if err := t.TryUpdate(key, value); err != nil {
		panic(fmt.Sprintf("Unhandled trie error: %v", err))
	}
}

// TryUpdate associates key with value in the trie. Subsequent calls to
// Get will return value. If value has length zero, any existing value
// is deleted from the trie and calls to Get will return nil.
//
// The value bytes must not be modified by the caller while they are
// stored in the trie.
//
// If a node was not found in the database, a MissingNodeError is returned.
func (t *Trie) TryUpdate(key, value []byte) error {
	k := keybytesToHex(key)
	if len(value) != 0 {
		_, n, err := t.insert(t.root, nil, k, ValueNode(value))
		if err != nil {
			return err
		}
		t.root = n
	} else {
		_, n, err := t.delete(t.root, nil, k)
		if err != nil {
			return err
		}
		t.root = n
	}
	return nil
}

func (t *Trie) insert(n node, prefix, key []byte, value node) (bool, node, error) {
	Count2 = Count2 + 1
	//若key为空，则该节点标记为valuenode并存储值
	if len(key) == 0 {
		if v, ok := n.(ValueNode); ok {
			return !bytes.Equal(v, value.(ValueNode)), value, nil
		}
		return true, value, nil
	}
	switch n := n.(type) {
	// 如果整个 key 与节点的 Key 相匹配，则只更新值并将节点标记为终止节点
	case *ShortNode:
		matchlen := prefixLen(key, n.Key)
		// If the whole key matches, keep this short node as is
		// and only update the value.
		if matchlen == len(n.Key) {
			dirty, nn, err := t.insert(n.Val, append(prefix, key[:matchlen]...), key[matchlen:], value)
			if !dirty || err != nil {
				return false, n, err
			}
			return true, &ShortNode{n.Key, nn, t.newFlag()}, nil
		}
		// 否则在 key 和节点的 Key 不同的位置处分支出一个新的节点
		// Otherwise branch out at the index where they differ.
		branch := &FullNode{flags: t.newFlag()}
		var err error
		_, branch.Children[n.Key[matchlen]], err = t.insert(nil, append(prefix, n.Key[:matchlen+1]...), n.Key[matchlen+1:], n.Val)
		if err != nil {
			return false, nil, err
		}
		_, branch.Children[key[matchlen]], err = t.insert(nil, append(prefix, key[:matchlen+1]...), key[matchlen+1:], value)
		if err != nil {
			return false, nil, err
		}
		// Replace this ShortNode with the branch if it occurs at index 0.
		// 如果节点的 Key 在匹配的位置是0，则将该节点替换为新的分支
		if matchlen == 0 {
			return true, branch, nil
		}
		// Otherwise, replace it with a short node leading up to the branch.
		//否则，用一个新的短节点替换该节点，该节点的 Key 为原来的 Key 的前 matchlen 个字符
		return true, &ShortNode{key[:matchlen], branch, t.newFlag()}, nil

	case *FullNode:
		// 如果已经存在该 key 的子节点，则递归进入该节点，并更新该节点的子节点
		dirty, nn, err := t.insert(n.Children[key[0]], append(prefix, key[0]), key[1:], value)
		if !dirty || err != nil {
			return false, n, err
		}
		n = n.copy()
		n.flags = t.newFlag()
		n.Children[key[0]] = nn
		return true, n, nil

	case nil:
		return true, &ShortNode{key, value, t.newFlag()}, nil

	case HashNode:
		// We've hit a part of the trie that isn't loaded yet. Load
		// the node and insert into it. This leaves all child nodes on
		// the path to the value in the trie.
		rn, err := t.resolveHash(n, prefix)
		if err != nil {
			return false, nil, err
		}
		dirty, nn, err := t.insert(rn, prefix, key, value)
		if !dirty || err != nil {
			return false, rn, err
		}
		return true, nn, nil

	default:
		panic(fmt.Sprintf("%T: invalid node: %v", n, n))
	}
}

// Delete removes any existing value for key from the trie.
func (t *Trie) Delete(key []byte) {
	if err := t.TryDelete(key); err != nil {
		panic(fmt.Sprintf("Unhandled trie error: %v", err))
	}
}

// TryDelete removes any existing value for key from the trie.
// If a node was not found in the database, a MissingNodeError is returned.
func (t *Trie) TryDelete(key []byte) error {
	k := keybytesToHex(key)
	_, n, err := t.delete(t.root, nil, k)
	if err != nil {
		return err
	}
	t.root = n
	return nil
}

// delete returns the new root of the trie with key deleted.
// It reduces the trie to minimal form by simplifying
// nodes on the way up after deleting recursively.
func (t *Trie) delete(n node, prefix, key []byte) (bool, node, error) {
	Count1 = Count1 + 1
	switch n := n.(type) {
	case *ShortNode:
		matchlen := prefixLen(key, n.Key)
		if matchlen < len(n.Key) {
			return false, n, nil // don't replace n on mismatch
		}
		if matchlen == len(key) {
			return true, nil, nil // remove n entirely for whole matches
		}
		// The key is longer than n.Key. Remove the remaining suffix
		// from the subtrie. Child can never be nil here since the
		// subtrie must contain at least two other values with keys
		// longer than n.Key.
		dirty, child, err := t.delete(n.Val, append(prefix, key[:len(n.Key)]...), key[len(n.Key):])
		if !dirty || err != nil {
			return false, n, err
		}
		switch child := child.(type) {
		case *ShortNode:
			// Deleting from the subtrie reduced it to another
			// short node. Merge the nodes to avoid creating a
			// ShortNode{..., ShortNode{...}}. Use concat (which
			// always creates a new slice) instead of append to
			// avoid modifying n.Key since it might be shared with
			// other nodes.
			return true, &ShortNode{concat(n.Key, child.Key...), child.Val, t.newFlag()}, nil
		default:
			return true, &ShortNode{n.Key, child, t.newFlag()}, nil
		}

	case *FullNode:
		dirty, nn, err := t.delete(n.Children[key[0]], append(prefix, key[0]), key[1:])
		if !dirty || err != nil {
			return false, n, err
		}
		n = n.copy()
		n.flags = t.newFlag()
		n.Children[key[0]] = nn

		// Check how many non-nil entries are left after deleting and
		// reduce the full node to a short node if only one entry is
		// left. Since n must've contained at least two children
		// before deletion (otherwise it would not be a full node) n
		// can never be reduced to nil.
		//
		// When the loop is done, pos contains the index of the single
		// value that is left in n or -2 if n contains at least two
		// values.
		pos := -1
		for i, cld := range &n.Children {
			if cld != nil {
				if pos == -1 {
					pos = i
				} else {
					pos = -2
					break
				}
			}
		}
		if pos >= 0 {
			if pos != 16 {
				// If the remaining entry is a short node, it replaces
				// n and its key gets the missing nibble tacked to the
				// front. This avoids creating an invalid
				// ShortNode{..., ShortNode{...}}.  Since the entry
				// might not be loaded yet, resolve it just for this
				// check.
				cnode, err := t.resolve(n.Children[pos], prefix)
				if err != nil {
					return false, nil, err
				}
				if cnode, ok := cnode.(*ShortNode); ok {
					k := append([]byte{byte(pos)}, cnode.Key...)
					return true, &ShortNode{k, cnode.Val, t.newFlag()}, nil
				}
			}
			// Otherwise, n is replaced by a one-nibble short node
			// containing the child.
			return true, &ShortNode{[]byte{byte(pos)}, n.Children[pos], t.newFlag()}, nil
		}
		// n still contains at least two values and cannot be reduced.
		return true, n, nil

	case ValueNode:
		return true, nil, nil

	case nil:
		return false, nil, nil

	case HashNode:
		// We've hit a part of the trie that isn't loaded yet. Load
		// the node and delete from it. This leaves all child nodes on
		// the path to the value in the trie.
		rn, err := t.resolveHash(n, prefix)
		if err != nil {
			return false, nil, err
		}
		dirty, nn, err := t.delete(rn, prefix, key)
		if !dirty || err != nil {
			return false, rn, err
		}
		return true, nn, nil

	default:
		panic(fmt.Sprintf("%T: invalid node: %v (%v)", n, n, key))
	}
}

func concat(s1 []byte, s2 ...byte) []byte {
	r := make([]byte, len(s1)+len(s2))
	copy(r, s1)
	copy(r[len(s1):], s2)
	return r
}

func (t *Trie) resolve(n node, prefix []byte) (node, error) {
	if n, ok := n.(HashNode); ok {
		return t.resolveHash(n, prefix)
	}
	return n, nil
}

func (t *Trie) resolveHash(n HashNode, prefix []byte) (node, error) {
	hash := BytesToHash(n)
	// fmt.Println("resolving", hash)
	val, err := t.db.Get(hash[:], nil)
	if val == nil || err != nil {
		error := &MissingNodeError{NodeHash: hash, Path: prefix}
		fmt.Println(error.Error())
		return nil, &MissingNodeError{NodeHash: hash, Path: prefix}
	} else {
		node := mustDecodeNode(hash[:], val)
		return node, nil
	}
}

// Hash returns the root hash of the trie. It does not write to the
// database and can be used even if the trie doesn't have one.
// func (t *Trie) Hash() Hash {
// 	hash, cached, _ := t.hashRoot(nil, nil)
// 	t.root = cached
// 	return common.BytesToHash(hash.(HashNode))
// }

// // Commit writes all nodes to the trie's memory database, tracking the internal
// // and external (for account tries) references.
func (t *Trie) Commit(onleaf LeafCallback) (root Hash, err error) {
	if t.db == nil {
		panic("commit called on trie with nil database")
	}
	hash, cached, err := t.hashRoot(t.db, onleaf)
	if err != nil {
		return Hash{}, err
	}
	t.root = cached
	return BytesToHash(hash.(HashNode)), nil //hash一定是一个HashNode
}

func (t *Trie) hashRoot(db *leveldb.DB, onleaf LeafCallback) (node, node, error) {
	if t.root == nil {
		return HashNode(emptyRoot.Bytes()), nil, nil
	}
	h := newHasher(onleaf)
	defer returnHasherToPool(h)
	return h.hash(t.root, db, true)
}

//func Validate(key, value []byte, proof [][]byte, stateRoot Hash, db *leveldb.DB) (bool, error) {
//	// create a new trie from the state root
//	t, err := NewTrie(stateRoot, db)
//	if err != nil {
//		return false, err
//	}
//
//	// retrieve the original value of the trie root node
//	origTrieHash := proof[0]
//	origTrieBytes := proof[1]
//	origTrieNode, err := DecodeNode(origTrieHash, origTrieBytes)
//	if err != nil {
//		return false, err
//	}
//
//	// set the original value of the trie root node in the new trie
//	err = db.Put(origTrieHash, origTrieBytes, nil)
//	if err != nil {
//		return false, err
//	}
//
//	// set the original node and value in the new trie
//	pathProof := proof[2:]
//	newNode := origTrieNode
//	for _, path := range pathProof {
//		newNode, err = applyProof(newNode, path)
//		if err != nil {
//			return false, err
//		}
//	}
//
//	t.Update(key, value)
//
//	// compare the new trie root hash to the original trie root hash
//	newTrieHash, err := t.Commit(nil)
//	if err != nil {
//		return false, err
//	}
//	return bytes.Equal(newTrieHash[:], origTrieHash), nil
//}

// ApplyProof applies the given proof to the given node and returns the resulting node.

//func applyProof(node node, proof []byte) (node, error) {
//	var err error
//	switch n := node.(type) {
//	case *ShortNode:
//		return n.Val, nil
//	case *FullNode:
//		n.Children[proofElement], err = applyProof(n.Children[proofElement], proof[1:])
//		if err != nil {
//			return nil, err
//		}
//		return n, nil
//	case *HashNode:
//		return nil, fmt.Errorf("cannot apply proof to hash node")
//	default:
//		return nil, fmt.Errorf("invalid node type: %T", node)
//	}
//	return node, nil
//}
