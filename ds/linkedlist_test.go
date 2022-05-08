package ds_test

import (
	"testing"

	"booleangate.org/ds/ds"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListAppend(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		t.Run("single", func(t *testing.T) {
			var ll ds.LinkedList[int]
			ll.Append(1)

			assert.Equal(t, 1, ll.Len())
			assert.Equal(t, []int{1}, ll.ToArray())
		})

		t.Run("multiple", func(t *testing.T) {
			var ll ds.LinkedList[int]
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
			var ll ds.LinkedList[int]

			_, err := ll.At(-1)
			assert.Error(t, err)

			_, err = ll.At(0)
			assert.Error(t, err)

			_, err = ll.At(1)
			assert.Error(t, err)
		})

		t.Run("not_empty", func(t *testing.T) {
			var ll ds.LinkedList[int]
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
