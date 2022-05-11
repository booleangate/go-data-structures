package gods_test

import (
	"testing"

	. "booleangate.org/gods"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListAppend(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		t.Run("single", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1)

			assert.Equal(t, 1, ll.Len())
			assert.Equal(t, []int{1}, ll.ToArray())
		})

		t.Run("multiple", func(t *testing.T) {
			var ll LinkedList[int]
			vals := []int{1, 2, 3}
			ll.Append(vals...)

			assert.Equal(t, len(vals), ll.Len())
			assert.Equal(t, vals, ll.ToArray())
		})
	})

	t.Run("non_empty", func(t *testing.T) {
		t.Run("single", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1)
			ll.Append(2)

			assert.Equal(t, 2, ll.Len())
			assert.Equal(t, []int{1, 2}, ll.ToArray())
		})

		t.Run("multiple", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1)
			ll.Append(2, 3)

			assert.Equal(t, 3, ll.Len())
			assert.Equal(t, []int{1, 2, 3}, ll.ToArray())
		})
	})
}

func TestLinkedListInsert(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		t.Run("single", func(t *testing.T) {
			var ll LinkedList[int]
			assert.NoError(t, ll.Insert(0, 1))

			assert.Equal(t, 1, ll.Len())
			assert.Equal(t, []int{1}, ll.ToArray())
		})

		t.Run("multiple", func(t *testing.T) {
			var ll LinkedList[int]
			vals := []int{1, 2, 3}
			assert.NoError(t, ll.Insert(0, vals...))

			assert.Equal(t, len(vals), ll.Len())
			assert.Equal(t, vals, ll.ToArray())
		})

		t.Run("out_of_bounds", func(t *testing.T) {
			var ll LinkedList[int]
			assert.Error(t, ll.Insert(1, 1))
			assert.Error(t, ll.Insert(-1, 1))
		})
	})

	t.Run("not_empty", func(t *testing.T) {
		t.Run("single_middle", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1, 2, 3)
			assert.NoError(t, ll.Insert(1, 4))

			assert.Equal(t, 4, ll.Len())
			assert.Equal(t, []int{1, 4, 2, 3}, ll.ToArray())
		})

		t.Run("multiple_middle", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1, 2, 3)
			assert.NoError(t, ll.Insert(1, 4, 5, 6))

			assert.Equal(t, 6, ll.Len())
			assert.Equal(t, []int{1, 4, 5, 6, 2, 3}, ll.ToArray())
		})
		t.Run("single_end", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1, 2, 3)
			assert.NoError(t, ll.Insert(ll.Len()-1, 4))

			assert.Equal(t, 4, ll.Len())
			assert.Equal(t, []int{1, 2, 4, 3}, ll.ToArray())
		})

		t.Run("multiple_end", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1, 2, 3)
			assert.NoError(t, ll.Insert(ll.Len()-1, 4, 5, 6))

			assert.Equal(t, 6, ll.Len())
			assert.Equal(t, []int{1, 2, 4, 5, 6, 3}, ll.ToArray())
		})
	})
}

func TestLinkedListDelete(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		t.Run("out_of_bounds", func(t *testing.T) {
			var ll LinkedList[int]
			_, err := ll.Delete(0)
			assert.Error(t, err)
			_, err = ll.Delete(1)
			assert.Error(t, err)
			_, err = ll.Delete(-1)
			assert.Error(t, err)
		})
	})

	t.Run("not_empty", func(t *testing.T) {
		t.Run("head", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1, 2, 3)
			val, err := ll.Delete(0)
			assert.NoError(t, err)

			assert.Equal(t, 2, ll.Len())
			assert.Equal(t, []int{2, 3}, ll.ToArray())
			assert.Equal(t, 1, val)
		})

		t.Run("middle", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1, 2, 3)
			val, err := ll.Delete(1)
			assert.NoError(t, err)

			assert.Equal(t, 2, ll.Len())
			assert.Equal(t, []int{1, 3}, ll.ToArray())
			assert.Equal(t, 2, val)
		})

		t.Run("tail", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1, 2, 3)
			val, err := ll.Delete(ll.Len() - 1)
			assert.NoError(t, err)

			assert.Equal(t, 2, ll.Len())
			assert.Equal(t, []int{1, 2}, ll.ToArray())
			assert.Equal(t, 3, val)
		})
	})
}

func TestLinkedListAt(t *testing.T) {
	t.Run("bounds_check", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			var ll LinkedList[int]

			_, err := ll.At(-1)
			assert.Error(t, err)

			_, err = ll.At(0)
			assert.Error(t, err)

			_, err = ll.At(1)
			assert.Error(t, err)
		})

		t.Run("not_empty", func(t *testing.T) {
			var ll LinkedList[int]
			ll.Append(1, 2, 3)

			v, err := ll.At(0)
			assert.NoError(t, err)
			assert.Equal(t, 1, v)

			v, err = ll.At(1)
			assert.NoError(t, err)
			assert.Equal(t, 2, v)

			v, err = ll.At(2)
			assert.NoError(t, err)
			assert.Equal(t, 3, v)
		})
	})
}

func TestLinkedListIterator(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var ll LinkedList[int]

		assert.False(t, ll.Iterator().Next(), "shouldn't iterator on an empty list")
	})

	t.Run("not_empty", func(t *testing.T) {
		var ll LinkedList[int]
		vals := []int{1, 2, 3}
		ll.Append(vals...)

		i := 0
		it := ll.Iterator()
		for it.Next() {
			val, ok := it.Value()
			assert.True(t, ok, "iterator value should be okay when not yet at the end")
			assert.Equal(t, vals[i], val)
			i++
		}

		_, ok := it.Value()
		assert.False(t, ok, "iterator value should not be okay after reaching the end")
		assert.False(t, it.Next(), "iterator should have no next")
		assert.Equal(t, i, 3, "iterator should have covered all elements in list")
	})
}

func TestLinkedListIteratorF(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var ll LinkedList[int]

		next := ll.IteratorF()
		_, ok := next()
		assert.False(t, ok, "Shouldn't iterator on an empty list")
	})

	t.Run("not_empty", func(t *testing.T) {
		var ll LinkedList[int]
		vals := []int{1, 2, 3}
		ll.Append(vals...)

		i := 0
		next := ll.IteratorF()
		for v, ok := next(); ok; v, ok = next() {
			assert.Equal(t, vals[i], v)
			i++
		}

		_, ok := next()
		assert.False(t, ok, "iterator should have no next")
		assert.Equal(t, i, 3, "iterator should have covered all elements in list")
	})
}
