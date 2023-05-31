package sort

import "math/rand"

func Quicksort(x Sortable) {

	shuffle(x)
	quicksort(x, 0, x.Len()-1)
}

func shuffle(x Sortable) {
	rand.Shuffle(x.Len(), func(i, j int) {
		x.Swap(i, j)
	})
}

// quicksort the subarray from x[lo] to x[hi]
func quicksort(x Sortable, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(x, lo, hi)
	quicksort(x, lo, j-1)
	quicksort(x, j+1, hi)
}

// partition the subarray x[lo...hi]
// so that x[lo...j-1 <= x[j+1...h]]
// and return the index j
func partition(x Sortable, lo, hi int) int {
	// let x[lo] be partition element
	i, j := lo+1, hi // exclude partition element

	for {
		// find item on lo to swap
		for i < hi && x.Less(i, lo) {
			i++
		}
		// find item on hi to swap
		for j > lo && !x.Less(j, lo) {
			j--
		}
		// check if pointers cross, if so don't swap
		if i >= j {
			break
		}
		// x[i] is greater than the partitioning element
		// and x[j] is smaller than the partitionign element.
		// swap them.
		x.Swap(i, j)
	}
	// put partitioning item at correct position x[j]
	x.Swap(lo, j)

	// now, x[lo...j-1] <= x[j] <= x[j+1...hi]
	// return index of partition element in its correct location
	return j
}

type Quick struct{}

func NewQuick() Sorter {
	return Quick{}
}

// Implements Sorter
func (s Quick) SortInts(x []int) {
	Quicksort(IntSortSlice(x))
}
func (s Quick) SortFloat64s(x []float64) {
	Quicksort(Float64SortSlice(x))
}
func (s Quick) SortStrings(x []string) {
	Quicksort(StringSortSlice(x))
}
