package sort

func InsertionSort(x Sortable) {
	// Sort x[] into increasing order
	n := x.Len()

	for i := 1; i < n; i++ {
		// insert x[i] among x[i-1], x[i-2], x[i-3]
		for j := i; j > 0 && x.Less(j, j-1); j-- {
			x.Swap(j, j-1)
		}
	}
}

type Insertion struct{}

func NewInsertion() Sorter {
	return Insertion{}
}

func (s Insertion) SortInts(x []int) {
	InsertionSort(IntSortSlice(x))
}
func (s Insertion) SortFloat64s(x []float64) {
	InsertionSort(Float64SortSlice(x))
}
func (s Insertion) SortStrings(x []string) {
	InsertionSort(StringSortSlice(x))
}
