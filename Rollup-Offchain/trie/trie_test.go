package trie

import (
	"bytes"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

// 请使用命令 go test -v trie_test.go  trie.go hash.go hasher.go bytes.go node.go encoding.go errors.go
func TestTrieSetPutandGet(t *testing.T) {
	db, err := leveldb.OpenFile("./testdata", nil)
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	root1 := "56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421" //64字节->ascii码->字节，emptyroot
	//b1 := []byte(root1)
	//v, _ := hex.Decode(b1, b1) //64字节的十六进制字符串被解码为32字节
	//fmt.Println(Hex2Bytes(root1))
	//HexToHash会将输入的root字符串转变为Hash(32字节类型)，若字符串以0x开头，则将0x去掉，若长度为奇数
	//则首项添0，若长度大于32，则选取后32个字节
	tr, err := NewTrie(HexToHash(root1), db)

	if err != nil {
		t.Error(err)
		return
	}
	var expGet = []byte("value")
	tr.Update([]byte("key"), expGet)
	tr.Update([]byte("kez"), []byte("error"))
	tr.Update([]byte("keyyy"), []byte("error"))
	tr.Update([]byte("keyyyyy"), []byte("error"))
	tr.Update([]byte("ke"), []byte("error"))
	tr.Commit(nil)
	//fmt.Println(db.Get([]byte("k1"), nil))
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Printf("key:%x,value:%x\n", iter.Key(), iter.Value())
	}
	iter.Release()
	//
	fmt.Println("分割")
	var toGet []byte
	toGet = tr.Get([]byte("key"))
	if !bytes.Equal(expGet, toGet) {
		t.Error("Put value is not equal to Getting value from memory..., expecting", expGet, "but,", toGet)
		return
	}

	var trHash Hash
	trHash, err = tr.Commit(nil)

	//trbytes := string(trHash[:])
	fmt.Println("trHash:", trHash[:])

	h := newHasher(nil)
	hash, node, _ := h.hash(tr.root, db, true)
	fmt.Println(hash, node)

	//hash := BytesToHash(hashnode)
	tr = nil
	tr, err = NewTrie(trHash, db)
	if err != nil {
		t.Error(err)
		return
	}

	toGet = tr.Get([]byte("key"))
	if !bytes.Equal(expGet, toGet) {
		t.Error("Put value is not equal to Getting value from db..., expecting", expGet, "but,", toGet)
		return
	}
}
