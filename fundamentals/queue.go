package fundamentals

type Queue struct {
	first *Node
	last  *Node
	n     int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue) Size() int {
	return q.n
}

func (q *Queue) Enqueue(item Item) {
	// add item to the end of the list
	oldLast := q.last
	q.last = newNode(item, nil)
	oldLast.next = q.last
	q.n++
}

func (q *Queue) Dequeue() Item {
	// remove item from the beginning of the list
	item := q.first.item
	q.first = q.first.next
	if q.IsEmpty() {
		q.last = nil
	}
	q.n--

	return item
}

func (q *Queue) Peek() Item {
	if q.IsEmpty() {
		panic("queue is empty")
	}

	return q.first.item
}

func (q *Queue) Iterator() Iterator {

	items := make([]interface{}, q.n)

	i := 0
	cur := q.first
	for i < q.n {
		items[i] = cur.item
		cur = cur.next
		i++
	}
	return Iterator(items)

}
