package bloom_filter

import (
	"golang.org/x/exp/constraints"
)

const MinFilterArraySize = 10

type HashFunc[K constraints.Ordered] func(K) int
type BloomFilter[K constraints.Ordered] struct {
	array    []int
	hashFunc []HashFunc[K]
}

func Set[K constraints.Ordered](self *BloomFilter[K], key K) {
	for _, f := range self.hashFunc {
		v := f(key) % len(self.array)
		self.array[v]++
	}
}

func IsHaving[K constraints.Ordered](self *BloomFilter[K], key K) bool {
	l := len(self.hashFunc)
	hit := 0
	for _, f := range self.hashFunc {
		v := f(key) % len(self.array)
		if self.array[v] != 0 {
			hit++
		}
	}
	if hit == l {
		return true
	}
	return false
}

func NewBloomFilter[K constraints.Ordered](size int, hashFunc []HashFunc[K]) *BloomFilter[K] {
	if size < 10 {
		size = 10
	}
	if len(hashFunc) == 0 {
		panic("No HashFunc found.\n")
	}
	return &BloomFilter[K]{
		array:    make([]int, size),
		hashFunc: hashFunc,
	}
}
