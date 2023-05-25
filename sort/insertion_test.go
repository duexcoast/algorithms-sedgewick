package sort_test

import (
	. "github.com/duexcoast/algorithms-sedgewick/sort"

	"testing"
)

func TestInsertionSortInts(t *testing.T) {
	data := ints
	x := IntSortSlice(data[0:])
	InsertionSort(x)
	if !IsSorted(x) {
		t.Errorf("sorting: %v", ints)
		t.Errorf("    got: %v", data)
	}

}

func TestInsertionSortFloat64s(t *testing.T) {
	data := float64s
	x := Float64SortSlice(data[0:])
	InsertionSort(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", float64s)
		t.Errorf("    got %v", data)
	}
}

func TestInsertionSortStrings(t *testing.T) {
	data := strings
	x := StringSortSlice(data[0:])
	InsertionSort(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", strings)
		t.Errorf("    got %v", data)
	}
}
