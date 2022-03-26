package hyperloglog_golang

import (
	"math"
)

// TODO: 並列数
func HyperLogLog(inputData []string, hashFunc func(string) uint32) float32 {
	// array := make([]float32, 0, len(inputData))
	// var max float32 = 0
	// 0000が2つ -> (2^4) * 2
	var max float32 = 0
	for _, v := range inputData {
		h := hashFunc(v)
		for pos := 32; pos >= 0; pos-- {
			// fmt.Printf("h >> pos %032b\n", h>>pos)
			// fmt.Printf("h        %032b\n", h)
			if (h>>(pos-1))&1 == 1 {
				if max < float32(pos) {
					max = float32(32 - pos)
				}
				break
			}
		}
	}
	ave := max / 1
	return float32(math.Pow(2, float64(ave)))
}
