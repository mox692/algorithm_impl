package util

import "testing"

func Test_PushAndGetNth(t *testing.T) {
	testCases := []int{0, 1, 2, 3}
	s := NewStack[int]()
	for i := 0; i < len(testCases); i++ {
		s.Push(testCases[i])
	}
	for i := 0; i < len(testCases); i++ {
		got := s.GetNth(i)
		if *got != testCases[len(testCases)-i-1] {
			t.Errorf("expect %d, but got %d\n", testCases[i], *got)
		}
	}
}
