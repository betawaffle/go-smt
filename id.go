package smt

import "hash"

const (
	// MaxDepth is the maximum number of layers a tree can hold.
	MaxDepth = HashSize * 8

	// MaxLayer is the maximum value of the layer byte in an ID.
	MaxLayer = MaxDepth - 1
	_        = uint8(MaxLayer) // ensure it fits
)

// ID represents the id of a node in a tree.
type ID [1 + HashSize]byte

func (id *ID) bit(layer uint8) uint8 {
	return id[HashSize-(layer>>3)] >> (layer & 7) & 1
}

func (id *ID) gotoSibling(i, j uint8) {
	id[i] ^= j
}

func (id *ID) gotoParent(i, j uint8) (left bool) {
	left = id[i]&j == 0
	id[0]++     // increment the layer
	id[i] &^= j // clear the child bit
	return left
}

func (id *ID) setLeaf(k []byte, h hash.Hash) {
	id[0] = 0 // leaves are at layer 0
	h.Write(k)
	h.Sum(id[1:1])
	h.Reset()
}
