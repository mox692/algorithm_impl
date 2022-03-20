package bloom_filter

import (
	"crypto/md5"
	"crypto/sha1"
)

func ShaFunc() HashFunc[string] {
	return func(hash string) int {
		s := sha1.New()
		defer s.Reset()
		s.Write([]byte(hash))
		b := s.Sum(nil)
		sum := 0
		for _, v := range b {
			sum += int(v)
		}
		return sum
	}
}

// GetMd5 - get encoded password with md5
func Md5Func() HashFunc[string] {
	return func(hash string) int {
		m := md5.New()
		defer m.Reset()
		m.Write([]byte(hash))
		b := m.Sum(nil)
		sum := 0
		for _, v := range b {
			sum += int(v)
		}
		return sum
	}
}
