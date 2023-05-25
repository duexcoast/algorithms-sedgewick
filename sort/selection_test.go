package sort_test

import (
	"testing"

	. "github.com/duexcoast/algorithms-sedgewick/sort"
)

func TestSelectionSortInts(t *testing.T) {
	data := ints
	x := IntSortSlice(data[0:])
	SelectionSort(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", ints)
		t.Errorf("    got %v", data)
	}
}

func TestSelectionSortFloat64s(t *testing.T) {
	data := float64s
	x := Float64SortSlice(data[0:])
	SelectionSort(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", float64s)
		t.Errorf("    got %v", data)
	}
}

func TestSelectionSortStrings(t *testing.T) {
	data := strings
	x := StringSortSlice(data[0:])
	SelectionSort(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", strings)
		t.Errorf("    got %v", data)
	}
}
