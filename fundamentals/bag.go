package fundamentals

type Bag[T any] struct {
	top *node[T]
	len int
}

func NewBag[T any]() *Bag[T] {
	return &Bag[T]{}
}

func (b *Bag[T]) IsEmpty() bool {
	return b.top == nil
}

func (b *Bag[T]) Size() int {
	return b.len
}

func (b *Bag[T]) Add(val T) {
	oldtop := b.top
	b.top = newNode(val)
	b.top.next = oldtop
	b.len++
}

func (b *Bag[T]) Peek() T {
	if b.IsEmpty() {
		panic("bag is empty")
	}
	return b.top.val

}

func (b *Bag[T]) Iterator() []T {
	items := make([]T, b.len)

	i := 0
	cur := b.top
	for i < b.len {
		items[i] = cur.val
		cur = cur.next
		i++
	}
	return items
}
