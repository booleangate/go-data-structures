package ds

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

func (l *LinkedList[T]) Append(vals ...T) {
	if len(vals) == 0 {
		return
	}

	chain := &llNode[T]{val: vals[0]}
	curr := chain
	valLen := len(vals)
	vals = vals[1:]
	for i := range vals {
		curr.next = &llNode[T]{val: vals[i]}
		curr = curr.next
	}

	if l.head == nil {
		l.head = chain
	} else {
		tail := l.head
		for tail.next != nil {
			tail = tail.next
		}
		tail.next = chain
	}
	l.len += valLen
}

func (l *LinkedList[T]) At(idx int) (T, error) {
	if idx < 0 || idx >= l.len {
		return zeroValue[T](), fmt.Errorf("index %d is out of bounds [0, %d)", idx, l.len)
	}

	var v T
	l.RangeUntil(func(i int, val T) (cont bool) {
		if i == idx {
			v = val
			return false
		}
		return true
	})

	return v, nil
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

func (l *LinkedList[T]) RangeUntil(cb func(idx int, val T) (cont bool)) {
	i := 0
	for n := l.head; n != nil; n = n.next {
		if cont := cb(i, n.val); !cont {
			break
		}
		i++
	}
}

func (l *LinkedList[T]) ToArray() []T {
	if l.head == nil {
		return nil
	}

	a := make([]T, 0, l.len)
	l.Range(func(i int, v T) {
		a = append(a, v)
	})

	return a
}
