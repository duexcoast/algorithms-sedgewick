package sort

func SelectionSort(x Sortable) {
	// array length
	n := x.Len()

	// exchange a[i] with smallest entry in a[i+1...n]
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if x.Less(j, min) {
				min = j
			}
		}
		x.Swap(i, min)
	}
}

type Selection struct{}

func NewSelection() Sorter {
	return Selection{}
}

// Implements Sorter
func (s Selection) SortInts(x []int) {
	SelectionSort(IntSortSlice(x))
}
func (s Selection) SortFloat64s(x []float64) {
	SelectionSort(Float64SortSlice(x))
}
func (s Selection) SortStrings(x []string) {
	SelectionSort(StringSortSlice(x))
}
