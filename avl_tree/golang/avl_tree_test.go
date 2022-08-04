package avl_tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// func Test_Bisic(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(1, "one")
// 	tr.Set(2, "two")
// 	tr.Delete(2)
// 	tr.Set(4, "four")

// 	// fmt.Println(flatten(tr.root))
// 	// tr.Set(2, "two")
// 	// tr.Set(3, "three")
// 	// tr.Set(4, "four")
// 	// tr.Set(5, "five")
// 	if got := *tr.Get(1); got != "one" {
// 		t.Errorf("expect %s, got %s\n", "one", got)
// 	}
// 	// if got := *tr.Get(2); got != "two" {
// 	// 	t.Errorf("expect %s, got %s\n", "two", got)
// 	// }
// 	// if got := *tr.Get(3); got != "three" {
// 	// 	t.Errorf("expect %s, got %s\n", "three", got)
// 	// }
// 	if got := *tr.Get(4); got != "four" {
// 		t.Errorf("expect %s, got %s\n", "four", got)
// 	}
// 	// if got := *tr.Get(5); got != "five" {
// 	// 	t.Errorf("expect %s, got %s\n", "zero", got)
// 	// }
// 	fmt.Println(flatten(tr.root))
// }

// func Test_Basic2(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(2, "one")
// 	tr.Set(3, "one")
// 	tr.Set(1, "one")
// 	fmt.Println(flatten(tr.root))
// 	tr.Delete(2)
// 	fmt.Println(flatten(tr.root))
// }

// func Test_Basic3(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(2, "one")
// 	tr.Set(1, "one")
// 	tr.Set(24, "one")
// 	tr.Set(3, "one")
// 	tr.Set(33, "one")
// 	tr.Set(44, "one")
// 	fmt.Println(flatten(tr.root))
// 	// fmt.Println(flatten(tr.root))
// 	// tr.Set(31, "one")

// 	// tr.Delete(2)
// }

// func Test_Basic4(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(10, "one")
// 	tr.Set(15, "two")
// 	tr.Set(6, "three")
// 	tr.Set(8, "four")
// 	tr.Set(5, "one")
// 	tr.Set(4, "one")
// }

// func Test_Basic5(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(10, "one")
// 	tr.Set(4, "two")
// 	tr.Set(16, "three")
// 	tr.Set(20, "four")
// 	tr.Set(13, "one")
// 	tr.Set(14, "one")
// 	fmt.Println("1aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// }

// func Test_Basic6(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(10, "one")
// 	tr.Set(4, "two")
// 	tr.Set(16, "three")
// 	tr.Set(20, "four")
// 	tr.Set(13, "one")
// 	// fmt.Println("aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// 	tr.Set(12, "one")
// 	fmt.Println("2aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// }

// func Test_Basic7(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(10, "one")
// 	tr.Set(4, "two")
// 	tr.Set(12, "three")
// 	tr.Set(2, "four")
// 	tr.Set(6, "one")
// 	// fmt.Println("aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// 	tr.Set(5, "one")
// 	fmt.Println("2aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// }

// func Test_Basic8(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(10, "one")
// 	tr.Set(4, "two")
// 	tr.Set(12, "three")
// 	tr.Set(2, "four")
// 	tr.Set(6, "one")
// 	// fmt.Println("aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// 	tr.Set(8, "one")
// 	fmt.Println("2aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// }

// func Test_Basic9(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(10, "one")
// 	tr.Set(4, "two")
// 	tr.Set(12, "three")
// 	tr.Set(2, "four")
// 	tr.Set(6, "one")
// 	// fmt.Println("aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// 	tr.Set(8, "one")
// 	fmt.Println("2aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// }

// func Test_Basic10(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(1, "one")
// 	fmt.Println("1 done")
// 	tr.Set(2, "two")
// 	fmt.Println("2 done")
// 	tr.Set(3, "three")
// 	fmt.Println("3 done")
// 	fmt.Println("aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// 	tr.Set(4, "four")
// 	fmt.Println("4 done")
// 	tr.Set(5, "one")
// 	fmt.Println("aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
// }

// func Test_flatten(t *testing.T) {
// }

// func Test_checkTree(t *testing.T) {
// 	tr := NewAvlTree()
// 	tr.Set(1, "one")
// 	tr.Set(2, "two")
// 	tr.Set(3, "three")
// 	tr.Set(4, "four")
// 	tr.Set(5, "one")
// 	checkTree(tr)
// }

func Test_random(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// s := rand.Perm(5)
	// 1 0 4 2 3
	s := []int{1, 2, 0, 4, 3}
	fmt.Println("aaaaaaaaaaaaa", s, "aaaaaaaaaaaaa")
	tr := NewAvlTree()
	for _, v := range s {
		tr.Set(v, fmt.Sprintf("%d", v))
	}
	for _, v := range s {
		got := *tr.Get(v)
		if expect := fmt.Sprintf("%d", v); got != expect {
			t.Errorf("expect %s, but got %s\ns:%+v", expect, got, s)
		}
	}
	fmt.Println(flatten(tr.root))
	checkTree(tr)
}
