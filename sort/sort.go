package sort

type Sortable interface {
	// the number of elements in the collection
	Len() int

	// Reports whether the element with index i
	// must sort before the element with index j
	Less(i, j int) bool

	// Swap the elements with indexes i and j
	Swap(i, j int)
}

type Sorter interface {
	SortInts(x []int)
	SortFloat64s(x []float64)
	SortStrings(x []string)
}

// Reports whether data is sorted
func IsSorted(data Sortable) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
func IsSortedInts(data []int) bool {
	return IsSorted(IntSortSlice(data[0:]))
}

// Convenience types for common cases

type (
	// Attaches the methods of Sortable to []int
	// sorting in increasing order
	IntSortSlice []int

	// Implements Sortable for a [] float64
	// sorting in increasing order
	// with NaN values ordered before other values
	Float64SortSlice []float64

	// Attaches methods of Sortable to []string
	// sorting in increasing order
	StringSortSlice []string
)

// Int SortSlice implements
func (x IntSortSlice) Len() int {
	return len(x)
}
func (x IntSortSlice) Less(i, j int) bool {
	return x[i] < x[j]
}
func (x IntSortSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// Float64SortSlice implements
func (x Float64SortSlice) Len() int {
	return len(x)
}

// Note that floating-point comparison by itself is not a transitive relation:
// it does not report a consistent ordering for not-a-number (NaN) values.
// This implementation places NaN values before others, by using:
//
// x[i] < x[j] || (math.IsNaN(x[i]) && !math.IsNaN(x[j]))
func (x Float64SortSlice) Less(i, j int) bool {
	return x[i] < x[j] || (isNaN(x[i]) && !isNaN(x[j]))
}
func (x Float64SortSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// A copy of math.IsNaN to avoid a dependency on the math package
func isNaN(f float64) bool {
	return f != f
}

// StringSortSlice implements
func (x StringSortSlice) Len() int {
	return len(x)
}
func (x StringSortSlice) Less(i, j int) bool {
	return x[i] < x[j]
}
func (x StringSortSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
