package binary_search_golang

import (
	"math/rand"
	"testing"
	"time"

	"github.com/mox692/tdata/array"
)

func BenchmarkBinarySearch(b *testing.B) {
	var len = 10000
	a := array.NewArrayInt(len)
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Hour.Nanoseconds())
		want := rand.Int() % len
		b.StartTimer()

		BinarySearch(a, want)

		b.StopTimer()
	}
}

type BinarySearchTestCase struct {
	input  []int
	want   int
	expect struct {
		have  bool
		index int
	}
}

func TestBinaruSearch(t *testing.T) {
	testCase := []BinarySearchTestCase{
		{
			input: []int{2, 3, 4, 5},
			want:  3,
			expect: struct {
				have  bool
				index int
			}{
				have:  true,
				index: 1,
			},
		},
		{
			input: []int{2, 3, 4, 5},
			want:  22,
			expect: struct {
				have  bool
				index int
			}{
				have:  false,
				index: 0,
			},
		},
		{
			input: []int{1},
			want:  1,
			expect: struct {
				have  bool
				index int
			}{
				have:  true,
				index: 0,
			},
		},
		{
			input: []int{},
			want:  1,
			expect: struct {
				have  bool
				index int
			}{
				have:  false,
				index: 0,
			},
		},
	}
	for _, v := range testCase {
		have, index := BinarySearch(v.input, v.want)
		if have != v.expect.have || index != v.expect.index {
			t.Errorf("expect: %+v, but got %+v, %+v", v.expect, have, index)
		}
	}
}
