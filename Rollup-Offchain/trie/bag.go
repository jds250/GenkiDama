package trie

type Bag struct {
	Proof [][]byte
	Key   []byte
	Value []byte
}

func (t *Trie) Getproof(key []byte) Bag {
	b := new(Bag)
	b.Proof = t.Prove(key)
	b.Key = key
	b.Value = t.Get(key)
	return *b
}

type Header struct {
	StateRoot Hash
}
