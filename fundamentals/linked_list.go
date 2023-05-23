package fundamentals

type node[T any] struct {
	val  T
	next *node[T]
}

func newNode[T any](val T) *node[T] {
	return &node[T]{
		val:  val,
		next: nil,
	}
}
