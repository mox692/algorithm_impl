package hyperloglog_golang

import (
	"math"
)

// 文字列の一覧(sliceを含む)を受け取り、そのユニークな要素数の推定値を返す.
// 上位ビットから0のビットがいくつ連続したかで推定.
func HyperLogLog(inputData []string, hashFunc func(string) uint32) float32 {
	var max float32 = 0
	for _, v := range inputData {
		h := hashFunc(v)
		for pos := 32; pos >= 1; pos-- {
			if (h>>(pos-1))&1 == 1 {
				if max < float32(32-pos) {
					max = float32(32 - pos)
				}
				break
			}
		}
	}

	// TODO: なんらかの方法で同じ入力に対して複数のhash関数を用意したい
	ave := max / 1
	return float32(math.Pow(2, float64(ave)))
}
