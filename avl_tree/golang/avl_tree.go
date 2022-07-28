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
	key    *int
	val    string
	parent *node
	l      *node
	r      *node
	lh     int
	rh     int
}

func newNode(parent *node, key *int, val string) *node {
	return &node{
		key:    key,
		val:    val,
		parent: parent,
		r:      nil,
		l:      nil,
		lh:     0,
		rh:     0,
	}
}

func newNodeDefault() *node {
	return &node{
		key:    nil,
		val:    "",
		parent: nil,
		r:      nil,
		l:      nil,
		lh:     0,
		rh:     0,
	}
}

func isBalanced(n *node) balanceType {
	if diff := n.lh - n.rh; diff >= 0 {
		// left heavy
		if diff >= 2 {
			lc := n.l
			if lc.r == nil && lc.l != nil {
				return UnBalancedLinearLeft
			} else if lc.r != nil && lc.l == nil {
				return UnBalancedCrookedLeft
			} else {
				log.Panicf("Invalid node, node: %+v\n", lc)
			}
		}
		return balanced
	} else {
		// right heavy
		if diff <= -2 {
			rc := n.r
			if rc.l == nil && rc.r != nil {
				return UnBalancedLinearRight
			} else if rc.r == nil && rc.l != nil {
				return UnBalancedCrookedRight
			} else {
				log.Panicf("Invalid node, node: %+v\n", rc)
			}
		}
		return balanced
	}
}

type direction int

const (
	left direction = iota
	right
)

type balanceType int

const (
	balanced balanceType = iota
	UnBalancedCrookedRight
	UnBalancedCrookedLeft
	UnBalancedLinearRight
	UnBalancedLinearLeft
)

// そのnodeがbalanceされているか.
// されていなかったら、バランス処理を施してreturnする.
// rootまで来たら(parentがなかったら)捜査は終了する
// from -> left: 0, right:1
func checkBalanceRec(n *node, from direction) {
	if from == left {
		n.lh++
	} else if from == right {
		n.rh++
	}
	parent := n.parent
	if b := isBalanced(n); b != balanced {
		// バランス処理.
		// バランス結果を反映するようにする.
		switch b {
		case UnBalancedLinearRight:
			ncopy := newNode(n, n.key, n.val)
			n.key = n.r.key
			n.val = n.r.val
			n.l = ncopy
			n.r = n.r.r
			if parent == nil {
				n.parent = nil
			} else {
				n.parent = parent
			}
			ncopy.parent = n
			n.r.parent = n
			n.rh = 1
			n.lh = 1
			n.r.rh = 0
			n.r.lh = 0
			n.l.rh = 0
			n.l.lh = 0
			return
		case UnBalancedLinearLeft:
			ncopy := newNode(n, n.key, n.val)
			n.key = n.l.key
			n.val = n.l.val
			n.r = ncopy
			n.l = n.l.l
			if parent == nil {
				n.parent = nil
			} else {
				n.parent = parent
			}
			ncopy.parent = n
			n.l.parent = n
			n.rh = 1
			n.lh = 1
			n.r.rh = 0
			n.r.lh = 0
			n.l.rh = 0
			n.l.lh = 0
			return
		case UnBalancedCrookedRight:
			//　　  1
			//       \
			//        3
			//       /
			//      2
			nr := n.r
			nrcopy := newNode(nr, nr.key, nr.val)
			nrcopy.parent = nr
			nr.key = nr.l.key
			nr.val = nr.l.val
			nr.r = nrcopy
			nr.l = nil
			//　　  1
			//       \
			//        2
			//         \
			//          3
			ncopy := newNode(n, n.key, n.val)
			n.key = n.r.key
			n.val = n.r.val
			n.l = ncopy
			n.r = n.r.r
			if parent == nil {
				n.parent = nil
			} else {
				n.parent = parent
			}
			ncopy.parent = n
			n.r.parent = n
			n.rh = 1
			n.lh = 1
			n.r.rh = 0
			n.r.lh = 0
			n.l.rh = 0
			n.l.lh = 0
		case UnBalancedCrookedLeft:
			//　　  3
			//    /
			//   1
			//    \
			//     2
			nl := n.l
			nlcopy := newNode(nl, nl.key, nl.val)
			nlcopy.parent = nl
			nl.key = nl.r.key
			nl.val = nl.r.val
			nl.l = nlcopy
			nl.r = nil
			//　　  3
			//    /
			//   2
			//  /
			// 1
			ncopy := newNode(n, n.key, n.val)
			n.key = n.l.key
			n.val = n.l.val
			n.r = ncopy
			n.l = n.l.l
			if parent == nil {
				n.parent = nil
			} else {
				n.parent = parent
			}
			ncopy.parent = n
			n.l.parent = n
			n.rh = 1
			n.lh = 1
			n.r.rh = 0
			n.r.lh = 0
			n.l.rh = 0
			n.l.lh = 0
		default:
			panic("not implement")
		}
	}
	if parent == nil {
		return
	}
	// そのnodeでbalanceしてるならparentをcheck
	if parent.r != nil {
		if parent.r.key == n.key {
			checkBalanceRec(parent, right)
			return
		}
		log.Fatal("errrrrrrrrr")
	}
	if parent.l != nil {
		if parent.l.key == n.key {
			checkBalanceRec(parent, left)
			return
		}
		log.Fatal("errrrrrrrrr")
	}
	return
}

func setRec(n *node, key *int, val string) {
	if *n.key < *key {
		if n.r == nil {
			n.r = newNode(n, key, val)
			checkBalanceRec(n, right)
			return
		}
		setRec(n.r, key, val)
		return
	}
	if *n.key > *key {
		if n.l == nil {
			n.l = newNode(n, key, val)
			checkBalanceRec(n, left)
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
	return fmt.Sprintf("([%d:%s, lh: %d, rh: %d], (%s, %s))", *n.key, n.val, n.lh, n.rh, l, r)
}
