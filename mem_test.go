package smt

import "testing"

func TestMem(t *testing.T) {
	testWriter(t, &memWriter{m: make(map[ID]Hash)})
}
