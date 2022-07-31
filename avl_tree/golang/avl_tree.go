package avl_tree

import (
	"fmt"
	"log"

	"github.com/mox692/algorithm_impl/avl_tree/golang/util"
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
	val, _ := getRec(t.root, key)
	return val
}

func getRec(n *node, key int) (*string, *node) {
	cur := n.key
	if key < *cur {
		if n.l == nil {
			return nil, n
		}
		return getRec(n.l, key)
	} else if key > *cur {
		if n.r == nil {
			return nil, n
		}
		return getRec(n.r, key)
	} else if key == *cur {
		return &n.val, n
	}
	return nil, n
}

func getRightChildRec(n *node) *node {
	if nr := n.r; nr != nil {
		return getRightChildRec(nr)
	}
	if p := n.parent; p != nil {
		p.r = nil
		// 親を再帰的に巡っていき、このnodeが消えることによるlh, rhを変更していく
	}
	return n
}

func (t avlTree) Delete(key int) {
	if t.root == nil {
		return
	}
	val, n := getRec(t.root, key)
	if val == nil {
		// keyに対応するデータがない
		return
	}
	// 削除は4ぱたーnn
	p := n.parent
	if n.l == nil && n.r == nil {
		if p == nil {
			// nを消すのではなく、初期化する
			n.key = nil
			n.val = ""
			n.l = nil
			n.r = nil
			n.rh = 0
			n.lh = 0
			return
		}
		if p.l != nil {
			if p.l.key == n.key {
				p.l = nil
				checkBalanceRec(p, util.NewStack[direction]().Push(left), delete)
				return
			}
		}
		if p.r != nil {
			if p.r.key == n.key {
				p.r = nil
				checkBalanceRec(p, util.NewStack[direction]().Push(right), delete)
				return
			}
		}
		panic("error happen")
	} else if nr := n.r; n.l == nil && nr != nil {
		if p == nil {
			// nを消すのではなく、初期化する
			n.key = nr.key
			n.val = nr.val
			n.l = nr.l
			n.r = nr.r
			n.rh--
			return
		}
		if p.l != nil {
			if p.l.key == n.key {
				p.l = nil
				checkBalanceRec(p, util.NewStack[direction]().Push(left), delete)
				return
			}
		}
		if p.r != nil {
			if p.r.key == n.key {
				p.r = nil
				checkBalanceRec(p, util.NewStack[direction]().Push(right), delete)
				return
			}
		}
		panic("error happen")
	} else if nl := n.l; nl != nil && n.r == nil {
		if p == nil {
			// nを消すのではなく、初期化する
			n.key = nr.key
			n.val = nr.val
			n.l = nl.l
			n.r = nl.r
			n.lh--
			return
		}
		if p.l != nil {
			if p.l.key == n.key {
				p.l = nil
				checkBalanceRec(p, util.NewStack[direction]().Push(left), delete)
				return
			}
		}
		if p.r != nil {
			if p.r.key == n.key {
				p.r = nil
				checkBalanceRec(p, util.NewStack[direction]().Push(right), delete)
				return
			}
		}
		panic("error happen")
	} else if n.l != nil && n.r != nil {
		// nのleftの一番大きいnodeを親に昇格させる
		var getBiggestNode func(*node) *node
		getBiggestNode = func(n *node) *node {
			if n.r == nil {
				return n
			}
			getBiggestNode(n.r)
			return nil
		}
		biggest := getBiggestNode(n.l)
		// 新しいbcに値を詰める
		n.key = biggest.key
		n.val = biggest.val
		// biggestを消しつつ、balancecheckして終了
		if bp := biggest.parent; bp.l.key == biggest.key {
			bp.l = nil
			fmt.Printf("bp: %+v\n", bp)
			checkBalanceRec(bp, util.NewStack[direction]().Push(left), delete)
			return
		} else if bp.r.key == biggest.key {
			bp.r = nil
			checkBalanceRec(bp, util.NewStack[direction]().Push(left), delete)
			return
		} else {
			panic("errrrrrrrrr")
		}
	}
}

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

func newNodeCopy(n *node) *node {
	if n == nil {
		return nil
	}
	return &node{
		key:    n.key,
		val:    n.val,
		parent: n.parent,
		r:      n.r,
		l:      n.l,
		lh:     n.lh,
		rh:     n.rh,
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

func isBalanced(n *node, path *util.Stack[direction]) balanceType {
	if diff := n.lh - n.rh; diff >= 0 {
		// left heavy
		if diff >= 2 {
			return dispatch(*path.GetNth(0), *path.GetNth(1))
		}
		return balanced
	} else {
		// right heavy
		if diff <= -2 {
			return dispatch(*path.GetNth(0), *path.GetNth(1))
		}
		return balanced
	}
}

func dispatch(first, second direction) balanceType {
	if first == right && second == right {
		return UnBalancedLinearRight
	} else if first == right && second == left {
		return UnBalancedCrookedRight
	} else if first == left && second == left {
		return UnBalancedLinearLeft
	} else if first == left && second == right {
		return UnBalancedCrookedLeft
	}
	panic("eerrrrrrrrrrr")
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
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

type checkType int

const (
	set    checkType = 1
	delete           = -1
)

// そのnodeがbalanceされているか.
// されていなかったら、バランス処理を施してreturnする.
// rootまで来たら(parentがなかったら)捜査は終了する
// from -> left: 0, right:1
func checkBalanceRec(n *node, path *util.Stack[direction], typ checkType) {
	if *path.GetNth(0) == left {
		n.lh += int(typ)
	} else if *path.GetNth(0) == right {
		n.rh += int(typ)
	}
	parent := n.parent
	if b := isBalanced(n, path); b != balanced {
		// バランス処理.
		// バランス結果を反映するようにする.
		fmt.Printf("come!! %+v\n", b)
		switch b {
		case UnBalancedLinearRight:
			ncopy := newNodeCopy(n)
			nrlcopy := newNodeCopy(n.r.l)

			// nの切り替え
			n.key = n.r.key
			n.val = n.r.val
			n.r = n.r.r

			// nl(元々n)の挿入
			ncopy.parent = n
			ncopy.l = n.l
			n.l = ncopy

			// nrlをnlにつける
			if nrlcopy != nil {
				nrlcopy.parent = n.l
			}
			n.l.r = nrlcopy

			// 高さ調整.新しいnと、nlが対象
			n.lh++
			n.rh--
			n.l.rh = max(n.l.r.rh, n.l.r.lh) + 1
			return
		case UnBalancedLinearLeft:
			ncopy := newNodeCopy(n)
			nlrcopy := newNodeCopy(n.l.r)

			// nの切り替え
			n.key = n.l.key
			n.val = n.l.val
			n.l = n.l.l

			// nl(元々n)の挿入
			ncopy.parent = n
			ncopy.r = n.r
			n.r = ncopy

			// nrlをnlにつける
			if nlrcopy != nil {
				nlrcopy.parent = n.r
			}
			n.r.l = nlrcopy

			// 高さ調整.新しいnと、nlが対象
			n.rh++
			n.lh--
			n.r.lh = max(n.r.l.rh, n.r.l.lh) + 1
			return
		case UnBalancedCrookedRight:
			//　　  1
			//       \
			//        3
			//       /
			//      2
			nr := n.r
			nrcopy := newNode(nr, nr.key, nr.val)
			nrcopy.parent = nr.l
			nrcopy.l = nil
			nr.key = nr.l.key
			nr.val = nr.l.val
			nr.r = nrcopy
			if nrlr := nr.l.r; nrlr != nil {
				nr.r.l = nrlr
			} else if nrll := nr.l.l; nrll != nil {
				nr.l = nrll
			} else {
				panic("wrong!!!")
			}
			//　　  1
			//       \
			//        2
			//         \
			//          3
			ncopy := newNodeCopy(n)
			nrlcopy := newNodeCopy(n.r.l)

			// nの切り替え
			n.key = n.r.key
			n.val = n.r.val
			n.r = n.r.r

			// nl(元々n)の挿入
			ncopy.parent = n
			ncopy.l = n.l
			n.l = ncopy

			// nrlをnlにつける
			if nrlcopy != nil {
				nrlcopy.parent = n.l
			}
			n.l.r = nrlcopy

			// 高さ調整.新しいnと、nlが対象
			n.lh++
			n.rh--
			n.l.rh = max(n.l.r.rh, n.l.r.lh) + 1
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
		if *parent.r.key == *n.key {
			checkBalanceRec(parent, path.Push(right), typ)
			return
		}
	}
	if parent.l != nil {
		if *parent.l.key == *n.key {
			checkBalanceRec(parent, path.Push(left), typ)
			return
		}
	}
}

func setRec(n *node, key *int, val string) {
	if *n.key < *key {
		if n.r == nil {
			n.r = newNode(n, key, val)
			// FIX: 新しいnodeに兄弟がいたら、親のheightだけを更新してreturnする
			if n.l != nil {
				n.rh++
				return
			}
			checkBalanceRec(n, util.NewStack[direction]().Push(right), set)
			return
		}
		setRec(n.r, key, val)
		return
	}
	if *n.key > *key {
		if n.l == nil {
			n.l = newNode(n, key, val)
			// FIX: 新しいnodeに兄弟がいたら、親のheightだけを更新してreturnする
			if n.r != nil {
				n.lh++
				return
			}
			checkBalanceRec(n, util.NewStack[direction]().Push(left), set)
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
