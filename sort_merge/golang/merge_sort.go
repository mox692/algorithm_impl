package main

import "fmt"

func mergeSort(v []int) []int {
	l := len(v)
	if l == 1 {
		return v
	}
	lv := v[0 : l>>1]
	rv := v[(l >> 1):l]
	lv = mergeSort(lv)
	rv = mergeSort(rv)

	fmt.Printf("lv: %+v\n", lv)
	fmt.Printf("rv: %+v\n", rv)
	res := make([]int, len(lv)+len(rv))
	fmt.Printf("res: %+v\n", res)
	rp := 0
	lp := 0
	for i := 0; i < len(res); i++ {
		if lp == len(lv) {
			fmt.Printf("conme1\n")
			rremain := rv[rp:]
			for j := lp + rp; j < len(res); j++ {
				res[j] = rremain[rp-1+j]
			}
			break
		} else if rp == len(rv) {
			fmt.Printf("conme2, res:%+v, lp+rp:%d, lp:%d, lv:%+v\n", res, lp+rp, lp, lv)
			lremain := lv[lp:]
			for j := lp + rp; j < len(res); j++ {
				res[j] = lremain[lp-1+j]
			}
			break
		}
		if lv[lp] <= rv[rp] {
			res[i] = lv[lp]
			lp++
		} else {
			res[i] = rv[rp]
			rp++
		}
	}
	fmt.Printf("res: %+v\n", res)
	return res
}
func main() {
	n := 0
	fmt.Scan(&n)
	input := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}
	res := mergeSort(input)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", res[i])
	}
}
