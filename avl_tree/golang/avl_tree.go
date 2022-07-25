package avl_tree

import (
	"fmt"
	"log"
)

type AVLTree interface {
	Set(key int, val string)
	Get(key int) *string
	Delete(key int)
}

type avlTree struct {
	root *node
}

func (t *avlTree) Set(key int, val string) {
	r := t.root
	if r.key == nil {
		r.key = &key
		r.val = val
		return
	}
	setRec(r, &key, val)

	// TODO: balance
	return
}

func (t avlTree) Get(key int) *string {
	if t.root == nil {
		return nil
	}
	return getRec(t.root, key)
}

func getRec(n *node, key int) *string {
	cur := n.key
	if key < *cur {
		if n.l == nil {
			return nil
		}
		return getRec(n.l, key)
	} else if key > *cur {
		if n.r == nil {
			return nil
		}
		return getRec(n.r, key)
	} else if key == *cur {
		return &n.val
	}
	return nil
}

func (t avlTree) Delete(key int) {}

func NewAvlTree() *avlTree {
	return &avlTree{
		root: newNodeDefault(),
	}
}

type node struct {
	key *int
	val string
	l   *node
	r   *node
}

func newNode(key *int, val string) *node {
	return &node{
		key: key,
		val: val,
		r:   nil,
		l:   nil,
	}
}
func newNodeDefault() *node {
	return &node{
		r: nil,
		l: nil,
	}
}

func setRec(n *node, key *int, val string) {
	if *n.key < *key {
		if n.r == nil {
			n.r = newNode(key, val)
			return
		}
		setRec(n.r, key, val)
		return
	}
	if *n.key > *key {
		if n.l == nil {
			n.l = newNode(key, val)
			return
		}
		setRec(n.l, key, val)
		return
	}
	if *n.key == *key {
		n.setKV(key, val)
		return
	}
	log.Fatalf("err")
}

func (n *node) setKV(key *int, val string) {
	n.key = key
	n.val = val
}

// helper
func flatten(n *node) string {
	if n == nil {
		return "NIL"
	}
	if n.l != nil && n.r != nil {
		if n.l.key == nil && n.r.key == nil {
			return string(n.val)
		}
	}
	l := "NIL"
	r := "NIL"
	if n.l != nil {
		if n.l.key != nil {
			l = flatten(n.l)
		}
	}
	if n.r != nil {
		if n.r.key != nil {
			r = flatten(n.r)
		}
	}
	return fmt.Sprintf("([%d:%s], (%s, %s))", *n.key, n.val, l, r)
}
