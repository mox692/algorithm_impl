package util

type Stack[T any] struct {
	head *elm[T]
}

type elm[T any] struct {
	val  T
	next *elm[T]
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func newElm[T any](e T) *elm[T] {
	n := new(elm[T])
	n.val = e
	return n
}

func newElmWithNext[T any](e T, next *elm[T]) *elm[T] {
	em := newElm(e)
	em.next = next
	return em
}

func (s *Stack[T]) Push(e T) *Stack[T] {
	old := s.head
	s.head = newElmWithNext(e, old)
	return s
}

func (s *Stack[T]) Pop() (*elm[T], *Stack[T]) {
	head := s.head
	s.head = head.next
	return head, s
}

func (s Stack[T]) Len() int {
	h := s.head
	if h == nil {
		return 0
	}
	l := 0
	for {
		l++
		if h.next != nil {
			h = h.next
		} else {
			break
		}
	}
	return l
}

// return nil if not exist
func (s Stack[T]) GetNth(n int) *T {
	if n < 0 {
		panic("negative number")
	}
	cur := s.head
	for i := 0; i < n; i++ {
		if cur == nil {
			return nil
		}
		cur = cur.next
	}
	return &cur.val
}
