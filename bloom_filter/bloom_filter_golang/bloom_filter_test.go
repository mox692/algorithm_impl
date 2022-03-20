package bloom_filter

import (
	"fmt"
	"testing"
)

func Test_BloomFilter(t *testing.T) {
	f := NewBloomFilter(10, []HashFunc[string]{ShaFunc()})

	Set(f, "")
	Set(f, "f")
	Set(f, "g")
	Set(f, "h")
	// must be true
	fmt.Printf("have f is %+v\n", IsHaving(f, "f"))
	fmt.Printf("have g is %+v\n", IsHaving(f, "g"))
	fmt.Printf("have h is %+v\n", IsHaving(f, "h"))
	// (it's actually false) may be true
	fmt.Printf("have i is %+v\n", IsHaving(f, "i"))
	fmt.Printf("%+v\n", f)
}

func Test_BloomFilter2(t *testing.T) {
	f := NewBloomFilter(10, []HashFunc[string]{ShaFunc(), Md5Func()})

	Set(f, "")
	Set(f, "f")
	Set(f, "g")
	Set(f, "h")
	// must be true
	fmt.Printf("have f is %+v\n", IsHaving(f, "f"))
	fmt.Printf("have g is %+v\n", IsHaving(f, "g"))
	fmt.Printf("have h is %+v\n", IsHaving(f, "h"))
	// (it's actually false) may be true
	fmt.Printf("have i is %+v\n", IsHaving(f, "i"))
	fmt.Printf("%+v\n", f)
}
