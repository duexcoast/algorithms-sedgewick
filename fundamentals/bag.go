package fundamentals

type Bag struct {
	top *Node
	n   int
}

func NewBag() *Bag {
	return &Bag{}
}

func (b *Bag) IsEmpty() bool {
	return b.top == nil
}

func (b *Bag) Size() int {
	return b.n
}

func (b *Bag) Add(item Item) {
	oldtop := b.top
	b.top = newNode(item, b.top)
	b.top.next = oldtop
	b.n++
}

func (b *Bag) Iterator() Iterator {
	items := make([]interface{}, b.n)

	i := 0
	cur := b.top
	for i < b.n {
		items[i] = cur.item
		cur = cur.next
		i++
	}
	return Iterator(items)
}
