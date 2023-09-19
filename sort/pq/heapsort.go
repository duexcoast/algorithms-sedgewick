package pq

type sortHeap []*elem

type elem struct {
	val      any
	index    int
	priority int
}

func (sh *sortHeap) Len() int {
	return len(*sh)
}

func (sh *sortHeap) Less(i, j int) bool {
	return (*sh)[i].priority > (*sh)[j].priority
}

func (sh *sortHeap) Swap(i, j int) {
	(*sh)[i], (*sh)[j] = (*sh)[j], (*sh)[i]
	(*sh)[i].index = i
	(*sh)[j].index = j
}
