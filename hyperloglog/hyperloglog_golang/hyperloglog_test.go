package hyperloglog_golang

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func Test_Hyperloglog(t *testing.T) {
	hashfunc := func(in string) uint32 {
		hash := md5.New()
		defer hash.Reset()
		hash.Write([]byte(in))
		s := hash.Sum(nil)
		s = s[:4]

		// 10000
		return uint32(s[0]) | uint32(s[1])<<8 | uint32(s[2])<<16 | uint32(s[3])<<24
		// fmt.Printf("aaaaaaaa %08b\n", s[0])
		// max := 0
		// for i := 0; i < len(in); i++ {
		// 	c := 0
		// 	for pos := 7; pos >= 0; pos-- {
		// 		fmt.Printf("aa %08b\n", s[0]>>pos)
		// 		if (s[0]>>pos)&1 == 1 {
		// 			break
		// 		}
		// 		c++
		// 	}
		// 	fmt.Printf("c is %d\n", c)
		// 	if max < c {
		// 		max = c
		// 	}
		// }
		// fmt.Printf("max is %d\n", max)
		// return 0
	}
	data := [][]string{
		{"0", "1", "a", "b", "c", "d", "e", "f", "g", "h", "i"},
		{"4fdsa", "xff3", "fdsai"},
		{"4a", "f3", "ai", "fsa,", ";fa", "cfdfa", "afd", "fdsa", "fasd", "fdsafdasfdfsd"},
	}
	for _, v := range data {
		fmt.Printf("res, %f\n", HyperLogLog(v, hashfunc))
	}
	// for _, v := range data {
	// }
}
