package hashtable

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	var keys = []string{"H", "A", "S", "H", "T", "A", "B", "L", "E"}
	var val = "v"
	hashTable := NewHashTable(4)
	for _, k := range keys {
		err := hashTable.Put(k, val)
		if err != nil {
			t.Error(err)
		}
	}

	for _, k := range keys {
		err, v := hashTable.Get(k)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Get key %s, val %s", k, v)
	}
}
