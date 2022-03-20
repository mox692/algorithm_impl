package bloom_filter

import (
	"fmt"
	"testing"
)

func Test_BloomFilter(t *testing.T) {
	f := NewBloomFilter(1000, []HashFunc[string]{ShaFunc()})

	fmt.Printf("have f is %+v\n", IsHaving(f, "f"))
	Set(f, "")
	Set(f, "f")
	Set(f, "g")
	Set(f, "h")
	fmt.Printf("have f is %+v\n", IsHaving(f, "f"))
	fmt.Printf("%+v\n", f)
}
