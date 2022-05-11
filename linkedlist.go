package gods

import "fmt"

// LinkedList implements a singly-linked list
type LinkedList[T any] struct {
	head *llNode[T]
	len  int
}

type llNode[T any] struct {
	next *llNode[T]
	val  T
}

func newNodeChain[T any](vals []T) (head, tail *llNode[T]) {
	head = &llNode[T]{val: vals[0]}
	tail = head
	vals = vals[1:]
	for i := range vals {
		tail.next = &llNode[T]{val: vals[i]}
		tail = tail.next
	}
	return head, tail
}

func (l *LinkedList[T]) Append(vals ...T) {
	if len(vals) == 0 {
		return
	}

	if l.head == nil {
		l.head, _ = newNodeChain(vals)
	} else {
		tail := l.head
		for tail.next != nil {
			tail = tail.next
		}
		tail.next, _ = newNodeChain(vals)
	}
	l.len += len(vals)
}

func (l *LinkedList[T]) Insert(idx int, vals ...T) error {
	if len(vals) == 0 {
		return nil
	}

	// Always allow insert at 0, even when empty
	if idx == 0 {
		head, tail := newNodeChain(vals)
		tail.next = l.head
		l.head = head
	} else {
		node, prev, err := l.node(idx)
		if err != nil {
			return err
		}

		head, tail := newNodeChain(vals)
		prev.next = head
		tail.next = node

	}
	l.len += len(vals)

	return nil
}

func (l *LinkedList[T]) At(idx int) (T, error) {
	node, _, err := l.node(idx)
	if err != nil {
		return zeroValue[T](), err
	}

	return node.val, nil
}

func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) Range(cb func(idx int, val T)) {
	i := 0
	for n := l.head; n != nil; n = n.next {
		cb(i, n.val)
		i++
	}
}

func (l *LinkedList[T]) ToArray() []T {
	if l.head == nil {
		return nil
	}

	a := make([]T, 0, l.len)
	l.Range(func(_ int, val T) {
		a = append(a, val)
	})

	return a
}

func (l *LinkedList[T]) Iterator() Iterator[T] {
	return newLLIterator(l.head)
}

func (l *LinkedList[T]) IteratorF() IteratorF[T] {
	curr := l.head

	return func() (val T, ok bool) {
		if curr == nil {
			return zeroValue[T](), false
		}

		val = curr.val
		curr = curr.next

		return val, true
	}
}

func (l *LinkedList[T]) node(idx int) (node, prev *llNode[T], err error) {
	if l.len == 0 {
		return nil, nil, fmt.Errorf("index %d is out of bounds, list is empty", idx)
	}
	if idx < 0 || idx >= l.len {
		return nil, nil, fmt.Errorf("index %d is out of bounds [0, %d)", idx, l.len)
	}

	i := 0
	node = l.head
	for node != nil {
		if i == idx {
			break
		}
		prev, node = node, node.next
		i++
	}

	return node, prev, nil
}

type llIterator[T any] struct {
	curr *llNode[T]
}

func newLLIterator[T any](head *llNode[T]) *llIterator[T] {
	// Insert a dummy node at start to allow the first call to Next to move to head.
	return &llIterator[T]{
		curr: &llNode[T]{next: head},
	}
}

func (it *llIterator[T]) Next() bool {
	if it.curr == nil {
		return false
	}

	it.curr = it.curr.next
	return it.curr != nil
}

func (it *llIterator[T]) Value() (val T, ok bool) {
	if it.curr == nil {
		return zeroValue[T](), false
	}
	return it.curr.val, true
}
