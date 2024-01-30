package gods

import "golang.org/x/exp/constraints"

func MergeSort[T constraints.Ordered](a []T) []T {
	if len(a) == 0 {
		return make([]T, 0)
	}

	return mergeSort(a, 0, len(a)-1)
}

func mergeSort[T constraints.Ordered](a []T, lo, hi int) []T {
	// We've reached the bottom of the call tree while dividing a.
	if lo == hi {
		return []T{a[lo]}
	}

	mid := (lo + hi) / 2
	left := mergeSort(a, lo, mid)
	right := mergeSort(a, mid+1, hi)
	return merge(left, right)
}

func merge[T constraints.Ordered](left, right []T) []T {
	sorted := make([]T, len(left)+len(right))
	l, r := 0, 0
	for i := range sorted {
		switch {
		case r >= len(right):
			sorted[i], l = left[l], l+1
		case l >= len(left):
			sorted[i], r = right[r], r+1
		case left[l] <= right[r]:
			sorted[i], l = left[l], l+1
		default:
			sorted[i], r = right[r], r+1
		}

	}
	return sorted
}
