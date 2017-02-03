package smt

import (
	"crypto"
	"crypto/sha512"
	"hash"
)

const (
	// HashFunc is the hash function used for keys and values.
	HashFunc = crypto.SHA512_256

	// HashSize is the number of bytes in the output of HashFunc.
	HashSize = sha512.Size256
)

// Hash is used to store the output of HashFunc.
type Hash [HashSize]byte

func (v *Hash) SetEmpty(layer uint8) {
	*v = emptyCache[layer]
}

func (v *Hash) SetBytes(b []byte) {
	copy(v[:], b[:HashSize])
}

func sumLeaf(sum *Hash, value []byte, h hash.Hash) {
	h.Write(prefixLeaf)
	h.Write(value)
	h.Sum(sum[:0])
	h.Reset()
}

func sumNode(sum, sib *Hash, swap bool, h hash.Hash) {
	h.Write(prefixNode)
	if swap {
		h.Write(sib[:])
		h.Write(sum[:])
	} else {
		h.Write(sum[:])
		h.Write(sib[:])
	}
	h.Sum(sum[:0])
	h.Reset()
}

var (
	emptyCache = [MaxDepth]Hash{}
	prefixLeaf = []byte{0}
	prefixNode = []byte{1}
)

func init() {
	h := HashFunc.New()
	x := &emptyCache[0]
	sumLeaf(x, nil, h)
	for layer := uint8(1); layer != 0; layer++ {
		y := &emptyCache[layer]
		*y = *x
		sumNode(y, y, false, h)
		x = y
	}
	return
}
