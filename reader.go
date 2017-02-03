package smt

// Reader is the interface that must be satisifed for a readable store.
type Reader interface {
	Get(*ID, *Hash) bool
}
