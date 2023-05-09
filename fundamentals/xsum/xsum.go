package xsum

import (
	"github.com/duexcoast/algorithms-sedgewick/fundamentals"
	"github.com/duexcoast/algorithms-sedgewick/sorting"
)

var (
	sorter sorting.Sorter = sorting.Merge{}

	binarySearch = fundamentals.NewBinarySearch()
)

// containsDuplicates function returns true if sorted array a contains any
// duplicated integers.
func containsDuplicates(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] == a[i-1] {
			return true
		}
	}

	return false
}
