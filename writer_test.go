package smt

import "testing"

var testData = map[string]string{
	"some key":       "some value",
	"some other key": "some other value",
	"a really cool":  "example map entry",
}

func testWriter(t *testing.T, store Writer) bool {
	var (
		err  error
		root = new(Hash)
		path = make([]byte, 0, HashSize*MaxDepth)
		hash = HashFunc.New()
	)
	for k, v := range testData {
		key, val := []byte(k), []byte(v)
		path, err = Put(store, key, val, path[:0], root, hash)
		if err != nil {
			t.Error(err)
			return false
		}
		if !Valid(path, key, val, hash, root) {
			t.Errorf("expected %x", root[:])
		}
	}
	return true
}
