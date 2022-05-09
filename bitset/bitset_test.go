package bitset

import (
	"fmt"
	"testing"
)

func TestBitSet_Add(t *testing.T) {
	bitset := NewBitSet(129)

	for i := 0; i < 200; i++ {
		set := bitset.Add(i)
		fmt.Printf("set %v \n", *set)
	}
}
