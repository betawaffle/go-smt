package smt

import (
	"os"
	"testing"

	"github.com/boltdb/bolt"
)

func TestBolt(t *testing.T) {
	d, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	defer os.Remove("test.db")

	fn := func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("test"))
		if err != nil {
			return err
		}
		testWriter(t, (*boltWriter)(b.Cursor()))
		return nil
	}
	if err := d.Update(fn); err != nil {
		t.Fatal(err)
	}
}
