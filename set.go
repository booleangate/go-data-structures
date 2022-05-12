package gods

type Hasher[K comparable, V any] func(V) K

type Set[K comparable, V any] struct {
	vals   map[K]V
	hasher Hasher[K, V]
}

func NewSet[T comparable]() Set[T, T] {
	// For a comparable, the hasher just returns the value
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

func (s *Set[K, V]) Delete(vals ...V) int {
	var deleted int
	for _, v := range vals {
		k := s.hasher(v)
		if _, ok := s.vals[k]; ok {
			delete(s.vals, k)
			deleted++
		}
	}
	return deleted
}

func (s *Set[K, V]) Clear() int {
	l := s.Len()
	s.vals = map[K]V{}
	return l
}

func (s *Set[K, V]) Has(val V) bool {
	k := s.hasher(val)
	_, ok := s.vals[k]
	return ok
}

func (s *Set[K, V]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set[K, V]) Len() int {
	return len(s.vals)
}

func (s *Set[K, V]) ToArray() []V {
	if s.IsEmpty() {
		return nil
	}

	vals := make([]V, 0, s.Len())
	for k := range s.vals {
		vals = append(vals, s.vals[k])
	}

	return vals
}

// Range ranges over the linked list and calls the callback, cb with the index and value of each node in the list.
func (s *Set[K, V]) Range(cb func(val V)) {
	for k := range s.vals {
		cb(s.vals[k])
	}
}

// Iterator returns a new iterator interface.
func (s *Set[K, V]) Iterator() Iterator[V] {
	if s.IsEmpty() {
		return setIterator[K, V]{}
	}

	keys := s.getKeys()
	// Add a sentinal at the beginning so we can call Next to start the iteration.
	keys.Insert(0, zeroValue[K]()) // nolint - we know that inserting at index 0 is always allowed

	return setIterator[K, V]{
		keys: keys,
		set:  s,
	}
}

// Iterator returns a new iterator function.
func (s *Set[K, V]) IteratorF() IteratorF[V] {
	keys := s.getKeys()

	return func() (val V, ok bool) {
		// If the set is mutated via deletion, we'll have keys that don't exist anymore.  Skip over those keys until we
		// find one that exists or we reach the end
		for !ok && !keys.IsEmpty() {
			key, _ := keys.Delete(0)
			val, ok = s.vals[key]
		}
		if ok {
			return val, true
		}

		return zeroValue[V](), false
	}
}

func (s *Set[K, V]) getKeys() LinkedList[K] {
	keys := LinkedList[K]{} // TODO: use a queue once that's written, it will have a better interface (push and pop)
	for k := range s.vals {
		keys.Append(k)
	}
	return keys
}

type setIterator[K comparable, V any] struct {
	keys LinkedList[K]
	set  *Set[K, V]
}

func (it setIterator[K, V]) Next() bool {
	if it.keys.IsEmpty() {
		return false
	}

	it.keys.Delete(0) // nolint
	return !it.keys.IsEmpty()
}
func (it setIterator[K, V]) Value() (val V, ok bool) {
	for !it.keys.IsEmpty() {
		// If the set is mutated via deletion, we'll have keys that don't exist anymore.  Skip over those keys until we
		// find one that exists or we reach the end
		for !ok && !it.keys.IsEmpty() {
			key, _ := it.keys.At(0)
			val, ok = it.set.vals[key]
			if ok {
				return val, true
			}
			// Delete the unknown key and try again
			it.keys.Delete(0) //nolint
		}
	}

	return zeroValue[V](), false
}
