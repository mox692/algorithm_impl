package avl_tree

import (
	"fmt"
	"testing"
)

func Test_Bisic(t *testing.T) {
	tr := NewAvlTree()
	tr.Set(1, "one")
	tr.Set(2, "one")
	fmt.Println(flatten(tr.root))
	tr.Delete(2)
	tr.Set(4, "one")

	// fmt.Println(flatten(tr.root))
	// tr.Set(2, "two")
	// tr.Set(3, "three")
	// tr.Set(4, "four")
	// tr.Set(5, "five")
	// if got := *tr.Get(1); got != "one" {
	// 	t.Errorf("expect %s, got %s\n", "one", got)
	// }
	// if got := *tr.Get(2); got != "two" {
	// 	t.Errorf("expect %s, got %s\n", "two", got)
	// }
	// if got := *tr.Get(3); got != "three" {
	// 	t.Errorf("expect %s, got %s\n", "three", got)
	// }
	// if got := *tr.Get(4); got != "four" {
	// 	t.Errorf("expect %s, got %s\n", "four", got)
	// }
	// if got := *tr.Get(5); got != "five" {
	// 	t.Errorf("expect %s, got %s\n", "zero", got)
	// }
}

func Test_flatten(t *testing.T) {
}
