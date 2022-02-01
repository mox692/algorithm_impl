package hashmap_golang

import (
	"testing"

	st "github.com/mox692/tdata/string"
)

func TestHashMapSet(t *testing.T) {
	maplen := 1000
	charLen := 15
	hashMapLen := 10000000

	numslice := make([]int, maplen)
	for i := 0; i < len(numslice); i++ {
		numslice[i] = i
	}
	keys := st.NewAsciiStringWithCharLenSlice(maplen, charLen)
	nums := numslice

	h := NewHashMapWithLen(hashMapLen)
	for i := 0; i < maplen; i++ {
		h.Set(keys[i], nums[i])
	}
	for i := 0; i < maplen; i++ {
		if got := h.Get(keys[i]); got != nums[i] {
			t.Errorf("something wrong...\nexpect: %d, got: %d", nums[i], got)
		}
	}
}
