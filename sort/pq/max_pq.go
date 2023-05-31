package pq

// MaxPQ using the heap interface
type Item interface {
}

type MaxPQ struct {
	items []Item
	n     int
}

//
// type MaxPQ struct {
// 	items []Item
// 	n     int
// 	BinaryHeap
// }
//
// func NewMaxPQN(n int) *MaxPQ {
// 	items := make([]Item, n+1) // items[0] is unused, so size is n+1
// 	heap := NewBinaryHeapBased1()
// 	return &MaxPQ{items, 0, heap}
// }
//
// func NewMaxPQ() *MaxPQ {
// 	items := make([]Item, 1)
// 	heap := NewBinaryHeapBased1()
// 	return &MaxPQ{items, 0, heap}
// }
//
// // Implement interface PriorityQueue
// func (pq *MaxPQ) Insert(item Item) {
// 	// double size of array if necessary
// 	if pq.n == len(pq.items)-1 {
// 		pq.resize(2 * len(pq.items))
// 	}
//
// 	// add item, and percolate it up to maintain heap invariant
// 	pq.n++
//
// 	// since we begin our indexing at items[1], the last key is items[n] not
// 	// items[n-1]
// 	lastLeaf := pq.GetLastLeafIndex(pq.n)
// 	pq.items[lastLeaf] = item
// 	pq.swim(lastLeaf, lastLeaf)
// }
//
// func (pq *MaxPQ) swim(child, max int) {
// 	for child > 0 {
// 		parent := pq.GetParentIndex(child, max)
// 		if pq.isHigherPriority(parent, child) {
// 			return
// 		}
// 		swap(pq.items, parent, child)
// 		child = parent
// 	}
// }
//
// // returns true if pq.items[i] is greater than than pq.items[j]
// func (pq *MaxPQ) isHigherPriority(i, j int) bool {
// 	i1 := pq.items[i]
// 	i2 := pq.items[j]
// 	return i1.CompareTo(i2) > 0
// }
//
// func (pq *MaxPQ) Delete() Item {
// 	if pq.IsEmpty() {
// 		panic("the priority queue is empty")
// 	}
//
// 	rootIdx := pq.GetRootIndex()
// 	lastLeaf := pq.GetLastLeafIndex(pq.n)
//
// 	item := pq.items[rootIdx] // retrieve the max key from top
//
// 	swap(pq.items, rootIdx, lastLeaf) // exchange with last item
//
// 	pq.sink(rootIdx, lastLeaf-1) // restore heap property
//
// 	pq.n-- //
//
// 	// resize array if number of items is less than 1/4th
// 	if pq.n > 0 && pq.n == (len(pq.items)-1)/4 {
// 		pq.resize(len(pq.items) / 2)
// 	}
//
// 	return item
// }
//
// func (pq *MaxPQ) sink(parent, max int) {
// 	higherPriorityChild := pq.getHighPriorityChild(parent, max)
//
// 	for
//
// 	// base case:
// 	// if the left and right child do not exist, stop sinking
// 	if higherPriorityChild == -1 {
// 		return
// 	}
//
// 	if pq.isHigherPriority(higherPriorityChild, parent) {
// 		swap(pq.items, higherPriorityChild, parent)
// 		pq.sink(higherPriorityChild, max)
// 	}
//
// }
//
// func (pq *MaxPQ) getHighPriorityChild(parent, max int) int {
// 	leftChild := pq.GetLeftChildIndex(parent, max)
// 	rightChild := pq.GetRightChildIndex(parent, max)
//
// 	if leftChild != -1 && rightChild != -1 {
// 		if pq.isHigherPriority(leftChild, rightChild) {
// 			return leftChild
// 		} else {
// 			return rightChild
// 		}
// 	} else if leftChild != -1 {
// 		return leftChild
// 	} else if rightChild != -1 {
// 		return rightChild
// 	} else {
// 		return -1
// 	}
// }
//
// func (pq *MaxPQ) IsEmpty() bool {
// 	return pq.n == 0
// }
//
// func (pq *MaxPQ) Size() int {
// 	return pq.n
// }
//
// func (pq *MaxPQ) resize(capacity int) {
// 	t := make([]Item, capacity)
// 	copy(t, pq.items)
// 	pq.items = t
// }
//
// func (pq *MaxPQ) GetItems() []Item {
// 	begin := pq.GetRootIndex()
// 	end := pq.GetLastLeafIndex(pq.n)
// 	return pq.items[begin : end+1]
// }
