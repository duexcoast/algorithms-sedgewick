package sort_test

import (
	"testing"

	. "github.com/duexcoast/algorithms-sedgewick/sort"
)

func TestQuicksortInts(t *testing.T) {
	data := ints
	x := IntSortSlice(data[0:])
	t.Logf("got: %v", x)

	Quicksort(x)
	t.Logf("got: %v", x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", ints)
		t.Errorf("    got %v", data)
	}

}
