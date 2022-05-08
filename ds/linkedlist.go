package ds

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

func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) Range(cb func(v T)) {
	for n := l.head; n != nil; n = n.next {
		cb(n.val)
	}
}

func (l *LinkedList[T]) ToArray() []T {
	if l.head == nil {
		return nil
	}

	a := make([]T, 0, l.len)
	l.Range(func(v T) {
		a = append(a, v)
	})

	return a
}

func (l *LinkedList[T]) Iterator() Iterator[T] {
	curr := l.head
	return func() (value T, hasNext bool) {
		v := curr.val
		if curr.next != nil {
			curr = curr.next
			return v, true
		}

		return v, false
	}
}
