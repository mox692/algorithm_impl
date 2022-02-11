package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	testCase := []struct {
		inputN   int
		inputArr []int
		expect   []int
	}{
		{
			inputN:   4,
			inputArr: []int{1, 5, 3, 2},
			expect:   []int{5, 2, 3, 1},
		},
		{
			inputN:   5,
			inputArr: []int{1, 9, 13, 25, 11},
			expect:   []int{25, 13, 9, 1, 11},
		},
	}

	for _, v := range testCase {
		got := heap(v.inputN, v.inputArr)
		if !sliceEq(got, v.expect) {
			t.Errorf("expect %+v, got %+v\n", v.expect, got)
		}
	}
}

func sliceEq(arg []int, cmp []int) bool {
	for i := 0; i < len(arg); i++ {
		if arg[i] != cmp[i] {
			return false
		}
	}
	return true
}
