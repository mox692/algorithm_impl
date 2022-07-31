package avl_tree

import (
	"fmt"
	"testing"
)

func Test_Bisic(t *testing.T) {
	tr := NewAvlTree()
	tr.Set(1, "one")
	tr.Set(2, "two")
	tr.Delete(2)
	tr.Set(4, "four")

	// fmt.Println(flatten(tr.root))
	// tr.Set(2, "two")
	// tr.Set(3, "three")
	// tr.Set(4, "four")
	// tr.Set(5, "five")
	if got := *tr.Get(1); got != "one" {
		t.Errorf("expect %s, got %s\n", "one", got)
	}
	// if got := *tr.Get(2); got != "two" {
	// 	t.Errorf("expect %s, got %s\n", "two", got)
	// }
	// if got := *tr.Get(3); got != "three" {
	// 	t.Errorf("expect %s, got %s\n", "three", got)
	// }
	if got := *tr.Get(4); got != "four" {
		t.Errorf("expect %s, got %s\n", "four", got)
	}
	// if got := *tr.Get(5); got != "five" {
	// 	t.Errorf("expect %s, got %s\n", "zero", got)
	// }
	fmt.Println(flatten(tr.root))
}

func Test_Basic2(t *testing.T) {
	tr := NewAvlTree()
	tr.Set(2, "one")
	tr.Set(3, "one")
	tr.Set(1, "one")
	fmt.Println(flatten(tr.root))
	tr.Delete(2)
	fmt.Println(flatten(tr.root))
}

func Test_Basic3(t *testing.T) {
	tr := NewAvlTree()
	tr.Set(2, "one")
	tr.Set(1, "one")
	tr.Set(24, "one")
	tr.Set(3, "one")
	tr.Set(33, "one")
	tr.Set(44, "one")
	fmt.Println(flatten(tr.root))
	// fmt.Println(flatten(tr.root))
	// tr.Set(31, "one")

	// tr.Delete(2)
}

func Test_Basic4(t *testing.T) {
	tr := NewAvlTree()
	tr.Set(10, "one")
	tr.Set(15, "two")
	tr.Set(6, "three")
	tr.Set(8, "four")
	tr.Set(5, "one")
	tr.Set(4, "one")
}

func Test_Basic5(t *testing.T) {
	tr := NewAvlTree()
	tr.Set(10, "one")
	tr.Set(4, "two")
	tr.Set(16, "three")
	tr.Set(20, "four")
	tr.Set(13, "one")
	fmt.Println("aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
	tr.Set(14, "one")
	fmt.Println("aaaaaaaaaaaaa", flatten(tr.root), "aaaaaaaaaaaaa")
}

func Test_flatten(t *testing.T) {
}
