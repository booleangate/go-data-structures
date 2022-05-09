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
