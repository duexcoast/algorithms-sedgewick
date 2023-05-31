package sort_test

import (
	"testing"

	. "github.com/duexcoast/algorithms-sedgewick/sort"
)

func TestMergesortInts(t *testing.T) {
	data := ints
	x := IntSortSlice(data[0:])
	sorter := NewMerge()
	sorter.SortInts(x)
	if !IsSorted(x) {
		t.Errorf("sorting: %v", ints)
		t.Errorf("    got: %v", x)
	}

}
