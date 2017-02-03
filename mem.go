package smt

type memReader memWriter

func (s *memReader) Get(id *ID, val *Hash) bool {
	return (*memWriter)(s).Get(id, val)
}

type memWriter struct {
	m map[ID]Hash
}

func (s *memWriter) Get(id *ID, val *Hash) (ok bool) {
	if val == nil {
		_, ok = s.m[*id]
		return
	}
	*val, ok = s.m[*id]
	if !ok {
		*val = emptyCache[id[0]]
	}
	return
}

func (s *memWriter) Put(id *ID, val *Hash) error {
	if val != nil {
		s.m[*id] = *val
	} else {
		delete(s.m, *id)
	}
	return nil
}

var _ Reader = &memReader{}
var _ Writer = &memWriter{}
