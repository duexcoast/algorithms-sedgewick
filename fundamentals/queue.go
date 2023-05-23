package fundamentals

type Queue[T any] struct {
	first *node[T]
	last  *node[T]
	len   int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue[T]) Size() int {
	return q.len
}

func (q *Queue[T]) Enqueue(val T) {
	// add item to the end of the list
	oldLast := q.last
	q.last = newNode(val)
	oldLast.next = q.last
	q.len++
}

func (q *Queue[T]) Dequeue() T {
	// remove item from the beginning of the list
	val := q.first.val
	q.first = q.first.next
	if q.IsEmpty() {
		q.last = nil
	}
	q.len--

	return val
}

func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		panic("queue is empty")
	}

	return q.first.val
}

func (q *Queue[T]) Iterator() []T {

	items := make([]T, q.len)

	i := 0
	cur := q.first
	for i < q.len {
		items[i] = cur.val
		cur = cur.next
		i++
	}
	return items

}
