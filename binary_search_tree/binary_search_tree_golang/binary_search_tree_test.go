package binary_search_tree

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	testCase := []struct {
		inputArr []int
		inputN   int
		expect   []int
	}{
		{
			inputArr: []int{2, 5, 1, 6},
			inputN:   4,
			expect:   []int{2, 1, 5, -1, -1, -1, 6},
		},
	}
	for _, v := range testCase {
		got := BinarySearchTree(v.inputN, v.inputArr)
		if !sliceEq(got, v.expect) {
			t.Errorf("got %+v, expect %+v\n", got, v.expect)
		}
	}
}

func sliceEq(s1, s2 []int) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
