package gods_test

import (
	"testing"

	. "booleangate.org/gods"
	"github.com/stretchr/testify/assert"
)

func TestSetAdd(t *testing.T) {
	t.Run("add_one", func(t *testing.T) {
		s := NewSet[int]()
		assert.Equal(t, 1, s.Add(1))
		assert.Equal(t, 1, s.Add(2))
		assert.ElementsMatch(t, []int{1, 2}, s.ToArray())
	})
	t.Run("add_many_unique", func(t *testing.T) {
		s := NewSet[int]()
		assert.Equal(t, 2, s.Add(1, 2))
		assert.Equal(t, 2, s.Add(3, 4))
		assert.ElementsMatch(t, []int{1, 2, 3, 4}, s.ToArray())
	})
	t.Run("add_many_nonunique", func(t *testing.T) {
		s := NewSet[int]()
		assert.Equal(t, 2, s.Add(1, 2))
		assert.Equal(t, 0, s.Add(1, 2))
		assert.ElementsMatch(t, []int{1, 2}, s.ToArray())
	})
}
