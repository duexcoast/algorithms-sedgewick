package fundamentals

type Item interface{}

type Node struct {
	item Item
	next *Node
}

func newNode(item Item, next *Node) *Node {
	return &Node{
		item: item,
		next: next,
	}
}
