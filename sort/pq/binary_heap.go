// package pq
//
// type BinaryHeap interface {
// 	GetLeftChildIndex(p, n int) int
// 	GetRightChildIndex(p, n int) int
// 	GetParentIndex(c, n int) int
// 	GetRootIndex() int
// 	GetLastLeafIndex(n int) int
// }
//
// // default 1-based indexing
// func NewBinaryHeap() BinaryHeap {
// 	return BinaryHeapBased1{}
// }
//
// func swap(a []Item, i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }
//
// type BinaryHeapBased1 struct{}
//
// func NewBinaryHeapBased1() BinaryHeapBased1 {
// 	return BinaryHeapBased1{}
// }
//
// // 1-based
// func (bh BinaryHeapBased1) GetLeftChildIndex(p, n int) int {
// 	leftChild := 2 * p
// 	if leftChild > n {
// 		return -1 // no valid left child
// 	}
//
// 	return leftChild
// }
//
// func (bh BinaryHeapBased1) GetRightChildIndex(p, n int) int {
// 	rightChild := 2*p + 1
// 	if rightChild > n {
// 		return -1 // no valid right child
// 	}
//
// 	return rightChild
// }
//
// func (bh BinaryHeapBased1) GetParentIndex(c, n int) int {
// 	if c <= 1 || c > n {
// 		return -1
// 	}
//
// 	return c / 2
// }
//
// func (bh BinaryHeapBased1) GetRootIndex() int {
// 	return 1
// }
//
// func (bh BinaryHeapBased1) GetLastLeafIndex(n int) int {
// 	return n
// }
