package ds

type Iterator[T any] func() (value T, hasNext bool)
