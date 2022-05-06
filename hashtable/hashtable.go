package hashtable

import (
	"errors"
	"fmt"
)

var (
	EmptyKey = errors.New("Key is empty")
	NotExist = errors.New("Key is not exist")
)

type HashTable struct {
	size    uint32 // 2^n
	buckets []*Bucket
}

func NewHashTable(size uint32) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([]*Bucket, size),
	}
}

func (ht *HashTable) Put(key, val string) error {
	if key == "" {
		return EmptyKey
	}
	index := ht.hashKey(key)

	bucket := ht.buckets[index]
	if bucket != nil {
		bucket.Put(key, val)
	} else {
		ht.buckets[index] = &Bucket{}
		bucket = ht.buckets[index]
		bucket.Put(key, val)
	}
	return nil
}

func (ht *HashTable) Get(key string) (error, string) {
	if key == "" {
		return EmptyKey, ""
	}
	index := ht.hashKey(key)
	bucket := ht.buckets[index]
	if bucket == nil {
		bucket = &Bucket{}
	}

	err, val := bucket.Get(key)
	if err != nil {
		return err, val
	}
	return nil, val
}

func (ht *HashTable) hashKey(key string) uint32 {
	// hash key
	hashCode := fnv32(key)
	// buckets index
	// when n is 2^n, h % n == h & (n - 1)
	index := hashCode & (ht.size - 1)
	fmt.Printf("key %s, hashcode %d, index %d \r\n", key, hashCode, index)
	return index
}

type Node struct {
	key   string
	value string
	next  *Node
}

type Bucket struct {
	head *Node
	len  int
}

func (b *Bucket) Add(key string, value string) {
	newHead := &Node{key: key, value: value, next: b.head}
	b.head = newHead
	b.len++
}

func (b *Bucket) Put(key string, value string) {
	isExist := false
	ptr := b.head
	for i := 0; i < b.len; i++ {
		// update
		if ptr.key == key {
			ptr.value = value
			isExist = true
		}
		ptr = ptr.next
	}

	if !isExist {
		// add
		b.Add(key, value)
	}
}

func (b *Bucket) Get(key string) (error, string) {
	ptr := b.head
	for i := 0; i < b.len; i++ {
		if ptr.key == key {
			return nil, ptr.value
		}
		ptr = ptr.next
	}
	return NotExist, ""
}

const prime32 = uint32(16777619)

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		v := uint32(key[i])
		hash ^= v
	}
	return hash
}
