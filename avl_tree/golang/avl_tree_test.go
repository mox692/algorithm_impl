package avl_tree

import (
	"fmt"
	"testing"
)

func Test_Bisic(t *testing.T) {
	tr := NewAvlTree()
	tr.Set(1, "three")
	tr.Set(2, "two")
	tr.Set(3, "one")
	tr.Set(4, "one")
	tr.Set(5, "one")
	tr.Set(6, "one")
	// tr.Set(4, "four")
	// tr.Set(0, "zero")
	if got := *tr.Get(3); got != "one" {
		t.Errorf("expect %s, got %s\n", "one", got)
	}
	if got := *tr.Get(2); got != "two" {
		t.Errorf("expect %s, got %s\n", "two", got)
	}
	if got := *tr.Get(1); got != "three" {
		t.Errorf("expect %s, got %s\n", "three", got)
	}
	fmt.Println(flatten(tr.root))
	// if got := *tr.Get(4); got != "four" {
	// 	t.Errorf("expect %s, got %s\n", "four", got)
	// }
	// if got := *tr.Get(0); got != "zero" {
	// 	t.Errorf("expect %s, got %s\n", "zero", got)
	// }
}

func Test_flatten(t *testing.T) {
	// tree := NewAvlTree()
	// tree.Set(3, "af")
	// tree.Set(2, "dsfasdasdf")
	// tree.Set(6, "afdasf")
	// fmt.Println(flatten(tree.root))
}