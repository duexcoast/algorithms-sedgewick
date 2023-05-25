package sort

import "reflect"

func MergesortBU(x Sortable) {
	n := x.Len()

	t := (reflect.TypeOf(x)).String()

	// convert type
	switch t {
	case "sorting.IntSortSlice":
		a := x.(IntSortSlice) // convert type
		sortIntsBU(a)         // merge sort on ints
	case "sorting.Float64SortSlice":
		a := make(Float64SortSlice, n)
		sortFloat64sBU(a)
	case "sorting.StringSortSlice":
		a := make(StringSortSlice, n)
		sortStringsBU(a)
	}
}

func sortIntsBU(x IntSortSlice) {
	n := x.Len()
	aux := make(IntSortSlice, n)

	for len := 1; len < n; len *= 2 { // len: subarray size
		for lo := 0; lo < n-len; lo = len * 2 { // lo: subarray index
			mid := lo + len - 1
			hi := min(lo+len*2-1, n-1)
			mergeIntsBU(x, aux, lo, mid, hi)
		}
	}
}

func mergeIntsBU(x, aux IntSortSlice, lo, mid, hi int) {
	// copy to aux
	copy(aux[lo:hi+1], x[lo:hi+1])

	// merge back to x
	i, j, k := lo, mid+1, lo

	for ; k <= hi; k++ {
		if i > mid {
			break
		} else if j > hi {
			x[k] = aux[i]
			i++
		} else if aux.Less(j, i) {
			x[k] = aux[j]
			j++
		} else {
			x[k] = aux[i]
			i++
		}
	}
}

func sortFloat64sBU(x Float64SortSlice) {
	n := x.Len()
	aux := make(Float64SortSlice, n)

	for len := 1; len < n; len *= 2 { // len: subarray size
		for lo := 0; lo < n-len; lo = len * 2 { // lo: subarray index
			mid := lo + len - 1
			hi := min(lo+len*2-1, n-1)
			mergeFloat64sBU(x, aux, lo, mid, hi)
		}
	}
}

func mergeFloat64sBU(x, aux Float64SortSlice, lo, mid, hi int) {
	// copy to aux
	copy(aux[lo:hi+1], x[lo:hi+1])

	// merge back to x
	i, j, k := lo, mid+1, lo

	for ; k <= hi; k++ {
		if i > mid {
			break
		} else if j > hi {
			x[k] = aux[i]
			i++
		} else if aux.Less(j, i) {
			x[k] = aux[j]
			j++
		} else {
			x[k] = aux[i]
			i++
		}
	}
}

func sortStringsBU(x StringSortSlice) {
	n := x.Len()
	aux := make(StringSortSlice, n)

	for len := 1; len < n; len *= 2 { // len: subarray size
		for lo := 0; lo < n-len; lo = len * 2 { // lo: subarray index
			mid := lo + len - 1
			hi := min(lo+len*2-1, n-1)
			mergeStringsBU(x, aux, lo, mid, hi)
		}
	}
}

func mergeStringsBU(x, aux StringSortSlice, lo, mid, hi int) {
	// copy to aux
	copy(aux[lo:hi+1], x[lo:hi+1])

	// merge back to x
	i, j, k := lo, mid+1, lo

	for ; k <= hi; k++ {
		if i > mid {
			break
		} else if j > hi {
			x[k] = aux[i]
			i++
		} else if aux.Less(j, i) {
			x[k] = aux[j]
			j++
		} else {
			x[k] = aux[i]
			i++
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type MergeBU struct{}

func NewMergeBU() Sorter {
	return MergeBU{}
}

// Implements Sorter
func (s MergeBU) SortInts(x []int) {
	MergesortBU(IntSortSlice(x))
}
func (s MergeBU) SortFloat64s(x []float64) {
	MergesortBU(Float64SortSlice(x))
}
func (s MergeBU) SortStrings(x []string) {
	MergesortBU(StringSortSlice(x))
}
