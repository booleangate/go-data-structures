package ds

type Iterator[T any] func() (value T, hasNext bool)

func zeroValue[T any]() T {
	var zv T
	return zv
}
