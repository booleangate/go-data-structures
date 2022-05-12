package gods

type Hasher[K comparable, V any] func(V) K

type Set[K comparable, V any] struct {
	vals   map[K]V
	hasher Hasher[K, V]
}

func NewSet[T comparable]() Set[T, T] {
	// For a comparable, the hasher just returns the value
	//
	return NewSetWithHasher(func(v T) T { return v })
}

func NewSetWithHasher[K comparable, V any](hasher Hasher[K, V]) Set[K, V] {
	return Set[K, V]{
		vals:   map[K]V{},
		hasher: hasher,
	}
}

func (s *Set[K, V]) Add(vals ...V) int {
	var added int
	for _, v := range vals {
		k := s.hasher(v)
		if _, ok := s.vals[k]; !ok {
			s.vals[k] = v
			added++
		}
	}
	return added
}

func (s *Set[K, V]) Delete(val ...V) int {
	panic("not implemented")
	return 0
}

func (s *Set[K, V]) Has(val V) bool {
	panic("not implemented")
	return false
}

func (s *Set[K, V]) Empty() bool {
	return s.Len() == 0
}

func (s *Set[K, V]) Len() int {
	return len(s.vals)
}

func (s *Set[K, V]) ToArray() []V {
	if s.Empty() {
		return nil
	}

	vals := make([]V, 0, len(s.vals))
	for k := range s.vals {
		vals = append(vals, s.vals[k])
	}

	return vals
}

// Range ranges over the linked list and calls the callback, cb with the index and value of each node in the list.
func (l *Set[K, V]) Range(cb func(val V)) {
	panic("not implemented")
}

// Iterator returns a new iterator interface.
func (l *Set[K, V]) Iterator() Iterator[V] {
	panic("not implemented")
	return nil
}

// Iterator returns a new iterator function.
func (l *Set[K, V]) IteratorF() IteratorF[V] {
	panic("not implemented")

	return func() (val V, ok bool) {
		return zeroValue[V](), false
	}
}
