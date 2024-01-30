package gods_test

import (
	"sort"
	"testing"

	. "booleangate.org/gods"

	"github.com/stretchr/testify/assert"
)

func FuzzMergeSort(f *testing.F) {
	f.Add([]byte(nil))
	f.Add([]byte(""))
	f.Add([]byte("90"))
	f.Add([]byte("v0v"))
	f.Add([]byte("lkj23rlkjzxvkljawvlkj23vklj2klj2v3lkj"))
	f.Fuzz(func(t *testing.T, input []byte) {
		act := MergeSort(input)
		exp := make([]byte, len(input))
		copy(exp, input)
		sort.Slice(exp, func(i, j int) bool {
			return exp[i] < exp[j]
		})
		assert.Equal(t, string(exp), string(act))
	})
}
