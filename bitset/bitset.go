package bitset

const (
	shift = 6    // 2^n = n of 64
	mask  = 0x3f // 2^6 - 1 = 63 = 0x3f
)

type BitSet struct {
	data []uint64
}

func NewBitSet(size int) *BitSet {
	return &BitSet{
		data: make([]uint64, size>>shift+1),
	}
}

func (set *BitSet) Add(n int) *BitSet {
	if n < 0 {
		return set
	}

	i := index(n)
	if i >= len(set.data) {
		newData := make([]uint64, i+1)
		copy(newData, set.data)
		set.data = newData
	}

	if set.data[i]&posVal(n) == 0 {
		set.data[i] |= posVal(n)
	}
	return set
}

func (set *BitSet) Clear(n int) *BitSet {
	if n < 0 {
		return set
	}

	i := index(n)
	if i >= len(set.data) {
		return set
	}

	if set.data[i]&posVal(n) != 0 {
		set.data[i] &^= posVal(n)
	}
	return set
}

func (set *BitSet) Contains(n int) bool {
	if n < 0 {
		return false
	}

	i := index(n)
	if i >= len(set.data) {
		return false
	}
	return set.data[i]&posVal(n) != 0
}

func index(n int) int {
	return n >> shift
}

func posVal(n int) uint64 {
	return 1 << uint(n&mask)
}
