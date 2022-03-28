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
		return uint32(s[0]) | uint32(s[1])<<8 | uint32(s[2])<<16 | uint32(s[3])<<24
	}
	data := [][]string{
		{"0", "1", "a", "b", "c", "d", "e", "f", "g", "h", "i"},
		{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16"},
		{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}, // 1
		{"b", "c", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}, // 3
		{"4fdsa", "xff3", "fdsai"},
		{"4a", "fdsafdasfdfsd"},
	}
	for _, v := range data {
		fmt.Printf("res, %f\n", HyperLogLog(v, hashfunc))
	}
}
