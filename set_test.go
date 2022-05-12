package gods_test

import (
	"math/rand"
	"testing"

	. "booleangate.org/gods"
	"github.com/stretchr/testify/assert"
)

func TestSetAdd(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		s := NewSet[int]()
		assert.Equal(t, 1, s.Add(1))
		assert.Equal(t, 1, s.Len())
		assert.Equal(t, 1, s.Add(2))
		assert.Equal(t, 2, s.Len())
		assert.ElementsMatch(t, []int{1, 2}, s.ToArray())
	})
	t.Run("many_unique", func(t *testing.T) {
		s := NewSet[int]()
		assert.Equal(t, 2, s.Add(1, 2))
		assert.Equal(t, 2, s.Len())
		assert.Equal(t, 2, s.Add(3, 4))
		assert.Equal(t, 4, s.Len())
		assert.ElementsMatch(t, []int{1, 2, 3, 4}, s.ToArray())
	})
	t.Run("many_nonunique", func(t *testing.T) {
		s := NewSet[int]()
		assert.Equal(t, 2, s.Add(1, 2))
		assert.Equal(t, 2, s.Len())
		assert.Equal(t, 0, s.Add(1, 2))
		assert.Equal(t, 2, s.Len())
		assert.ElementsMatch(t, []int{1, 2}, s.ToArray())
	})
}

func TestSetDelete(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(1, 2, 3, 4, 5)

		assert.Equal(t, 1, s.Delete(1))
		assert.Equal(t, 4, s.Len())
		assert.ElementsMatch(t, []int{2, 3, 4, 5}, s.ToArray())
	})
	t.Run("many_unique", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(1, 2, 3, 4, 5)

		assert.Equal(t, 3, s.Delete(1, 2, 3))
		assert.Equal(t, 2, s.Len())
		assert.ElementsMatch(t, []int{4, 5}, s.ToArray())
	})
	t.Run("many_nonunique", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(1, 2, 3, 4, 5)

		assert.Equal(t, 1, s.Delete(1, 1, 1))
		assert.Equal(t, 4, s.Len())
		assert.ElementsMatch(t, []int{2, 3, 4, 5}, s.ToArray())
	})
	t.Run("all", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(1, 2, 3, 4, 5)

		assert.Equal(t, 5, s.Delete(1, 2, 3, 4, 5))
		assert.Equal(t, 0, s.Len())
		assert.True(t, s.IsEmpty())
		var nilSlice []int
		assert.ElementsMatch(t, nilSlice, s.ToArray())
	})
}

func TestSetClear(t *testing.T) {
	var nilSlice []int

	t.Run("when_empty", func(t *testing.T) {
		s := NewSet[int]()

		assert.True(t, s.IsEmpty())
		assert.Equal(t, 0, s.Clear())
		assert.True(t, s.IsEmpty())
		assert.ElementsMatch(t, nilSlice, s.ToArray())
	})
	t.Run("when_not_empty", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(1, 2, 3, 4, 5)

		assert.False(t, s.IsEmpty())
		assert.Equal(t, 5, s.Clear())
		assert.True(t, s.IsEmpty())
		assert.ElementsMatch(t, nilSlice, s.ToArray())
	})
}

func TestSetHas(t *testing.T) {
	t.Run("when_empty", func(t *testing.T) {
		s := NewSet[int]()

		// Fuzzed
		for i := 0; i < 100; i++ {
			assert.False(t, s.Has(int(rand.Int31())))
		}
	})
	t.Run("not_empty_and_present", func(t *testing.T) {
		s := NewSet[int]()
		vals := []int{1, 2, 3, 4, 5}
		s.Add(vals...)

		for _, v := range vals {
			assert.True(t, s.Has(v))
		}
	})
	t.Run("not_empty_and_missing", func(t *testing.T) {
		s := NewSet[int]()
		vals := []int{1, 2, 3, 4, 5}
		s.Add(vals...)

		for _, v := range vals {
			assert.False(t, s.Has(-v))
		}
	})
}
