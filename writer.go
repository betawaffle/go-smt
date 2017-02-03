package smt

import "hash"

// Writer is the interface that must be satisfied for a writable store.
type Writer interface {
	Reader
	Put(*ID, *Hash) error
}

func Put(tree Writer, key, val, path []byte, root *Hash, hash hash.Hash) ([]byte, error) {
	if hash == nil {
		hash = HashFunc.New()
	}
	if root == nil {
		root = new(Hash)
	}

	id := ID{}
	id.setLeaf(key, hash)

	// compute and store the leaf hash
	sumLeaf(root, val, hash)
	if err := tree.Put(&id, root); err != nil {
		return nil, err
	}

	// walk up to the root
	for i, sibling := uint8(HashSize), new(Hash); i > 0; i-- {
		for j := uint8(1); j != 0; j <<= 1 {
			// fetch the sibling hash and add it to the path
			if id.gotoSibling(i, j); tree.Get(&id, sibling) {
				path = append(append(path, id[0]), sibling[:]...)
			}

			// compute and store the parent hash
			sumNode(root, sibling, id.gotoParent(i, j), hash)
			if err := tree.Put(&id, root); err != nil {
				return nil, err
			}

			if id[0] == MaxLayer {
				return path, nil
			}
		}
	}
	panic("unreachable")
}
