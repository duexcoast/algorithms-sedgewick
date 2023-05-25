package sort

import "reflect"

// Merging: combining two ordered arrays to make one larger ordered array.
// This opertation immeditaely lends itself to a simple recursive sort method
// known as mergesort: to sort an array, divide it into two halves, sort the
// two halves (recursively), and then merge the results.
//
// Mergesort guarantees to sort an array of N items in time proportional to NlogN,
// no matter what the input. Its prime disadvantage is that it uses extra space
// proportional to N
//
// Mergesort
// Top-down mergesort. It is one of the best-known examples of utility of the
// divide-and-conquer paradigm for efficient algorithm design.
func Mergesort(x Sortable) {
	n := x.Len()

	t := (reflect.TypeOf(x)).String()

	// convert type
	switch t {
	case "sorting.IntSortSlice":
		a := x.(IntSortSlice)        // convert type
		aux := make(IntSortSlice, n) // allocate space just once
		sortInts(a, aux, 0, n-1)     // merge sort on ints
	case "sorting.Float64SortSlice":
		a := x.(Float64SortSlice)
		aux := make(Float64SortSlice, n)
		sortFloat64s(a, aux, 0, n-1)
	case "sorting.StringSortSlice":
		a := x.(StringSortSlice)
		aux := make(StringSortSlice, n)
		sortStrings(a, aux, 0, n-1)
	}
}

// mergesort x[lo...hi] using auxiliary array aux[lo...hi]
func sortInts(x, aux IntSortSlice, lo, hi int) {
	if hi <= lo { // list has one record
		return
	}
	mid := lo + (hi-lo)/2          // Select midpoint
	sortInts(x, aux, lo, mid)      // Sort left half
	sortInts(x, aux, mid+1, hi)    // Sort right half
	mergeInts(x, aux, lo, mid, hi) // Merge results
}

func sortFloat64s(x, aux Float64SortSlice, lo, hi int) {
}

func sortStrings(x, aux StringSortSlice, lo, hi int) {
}

func mergeInts(x, aux IntSortSlice, lo, mid, hi int) {
	// Copy x[] to aux[]
	copy(aux, x)

	// Merge back to x[]
	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			// left is exhausted, merge from right
			x[k] = aux[j]
			j++
		} else if j > hi {
			// right is exhausted, merge from left
			x[k] = aux[i]
			i++
		} else if aux.Less(j, i) {
			// right is less than left, merge from right
			x[k] = aux[j]
		} else {
			// left is less than right, merge from left
			x[k] = aux[i]
		}
	}
}

type Merge struct{}

func NewMerge() Sorter {
	return Merge{}
}

// Implements Sorter
func (s Merge) SortInts(x []int) {
	Mergesort(IntSortSlice(x))
}
func (s Merge) SortFloat64s(x []float64) {
	Mergesort(Float64SortSlice(x))
}
func (s Merge) SortStrings(x []string) {
	Mergesort(StringSortSlice(x))
}
