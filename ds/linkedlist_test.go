package ds_test

import (
	"testing"

	"booleangate.org/ds/ds"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListAppend(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		t.Run("single", func(t *testing.T) {
			ll := ds.LinkedList[int]{}
			ll.Append(1)

			assert.Equal(t, 1, ll.Len())
			assert.Equal(t, []int{1}, ll.ToArray())
		})

		t.Run("multiple", func(t *testing.T) {
			ll := ds.LinkedList[int]{}
			vals := []int{1, 2, 3}
			ll.Append(vals...)

			assert.Equal(t, len(vals), ll.Len())
			assert.Equal(t, vals, ll.ToArray())
		})
	})
}
