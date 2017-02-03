package smt

import (
	"bytes"

	"github.com/boltdb/bolt"
)

type boltReader boltWriter

func (s *boltReader) Get(id *ID, val *Hash) bool {
	return (*boltWriter)(s).Get(id, val)
}

type boltWriter bolt.Cursor

func (s *boltWriter) Get(id *ID, val *Hash) bool {
	k, v := (*bolt.Cursor)(s).Seek(id[:])
	if k == nil || !bytes.Equal(k, id[:]) || len(v) < HashSize {
		if val != nil {
			*val = emptyCache[id[0]]
		}
		return false
	}
	if val != nil {
		copy(val[:], v)
	}
	return true
}

func (s *boltWriter) Put(id *ID, val *Hash) error {
	bucket := (*bolt.Cursor)(s).Bucket()
	if val == nil {
		return bucket.Delete(id[:])
	}
	copy := *val
	return bucket.Put(id[:], copy[:])
}

var _ Reader = &boltReader{}
var _ Writer = &boltWriter{}
