// package pq
//
// type PriorityQueue interface {
// 	Insert(val Item)
// 	Delete() Item
// 	IsEmpty() bool
// 	Size() int
// }
//
// // Index priority queue.
// // In many applications, it makes sense to allow client to refer to items that are already on the priority queue. One
// // easy way to do so is to associate an integer with each item.
// type IndexPQ interface {
// }
//
// type Item interface {
// 	CompareTo(x Item) int
// }
//
// type StringItem string
//
// func (s StringItem) CompareTo(x Item) int {
// 	ss := x.(StringItem)
// 	if s < ss {
// 		return -1
// 	} else if s > ss {
// 		return 1
// 	} else {
// 		return 0
// 	}
// }
//
// type IntItem int
//
// func (i IntItem) CompareTo(x Item) int {
// 	ii := x.(IntItem)
// 	if i < ii {
// 		return -1
// 	} else if i > ii {
// 		return 1
// 	} else {
// 		return 0
// 	}
// }
