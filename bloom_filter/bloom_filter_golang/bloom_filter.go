package bloom_filter

import (
	"crypto/sha1"

	"golang.org/x/exp/constraints"
)

const MinFilterArraySize = 10

func ShaFunc() HashFunc[string] {
	return func(hash string) int {
		s := sha1.New()
		s.Write([]byte(hash))
		b := s.Sum(nil)
		sum := 0
		for _, v := range b {
			sum += int(v)
		}
		return sum
	}
}

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
