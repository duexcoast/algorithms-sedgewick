package sort_test

import (
	"testing"

	. "github.com/duexcoast/algorithms-sedgewick/sort"
)

// go test -v -run="none" -bench="." -benchtime="3s"

var a = [...]int{9, 10, 0, 7, 8, 4, 3, 6, 2, 1, 5, 5, -99}

func BenchmarkSelection(b *testing.B) {
	sorter := NewSelection()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sorter.SortInts(a[0:])
	}

}
func BenchmarkInsertion(b *testing.B) {
	sorter := NewInsertion()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sorter.SortInts(a[0:])
	}

}
func BenchmarkShell(b *testing.B) {
	sorter := NewShell()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sorter.SortInts(a[0:])
	}

}
func BenchmarkMerge(b *testing.B) {
	sorter := NewMerge()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sorter.SortInts(a[0:])
	}

}
