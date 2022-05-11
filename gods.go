// gods provides GO Data Structures
package gods

// IteratorF defines a functional iterator.
type IteratorF[T any] func() (value T, ok bool)

// Iterator defines a struct iterator.
type Iterator[T any] interface {
	Next() bool
	Value() (val T, ok bool)
}

func zeroValue[T any]() T {
	var zv T
	return zv
}
