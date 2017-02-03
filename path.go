package smt

import "hash"

func Valid(path, key, val []byte, hash hash.Hash, expect *Hash) bool {
	actual := new(Hash)
	return sumPath(actual, path, key, val, hash) && *actual == *expect
}

func popNode(dst *Hash, path []byte, layer uint8) []byte {
	const (
		nodeSize = 1 + HashSize
	)
	if len(path) < nodeSize || layer < path[0] {
		dst.SetEmpty(layer)
		return path
	}
	if layer > path[0] {
		return nil // invalid path
	}
	dst.SetBytes(path[1:nodeSize]) // copy(dst[:], p[1:nodeSize])
	return path[nodeSize:]
}

func sumPath(sum *Hash, path, key, val []byte, h hash.Hash) bool {
	id := new(ID)
	id.setLeaf(key, h)
	sumLeaf(sum, val, h)
	for i, sibling := 0, new(Hash); i < MaxLayer; i++ {
		if path = popNode(sibling, path, uint8(i)); path == nil {
			return false
		}
		sumNode(sum, sibling, id.bit(uint8(i)) != 0, h)
	}
	return true
}
