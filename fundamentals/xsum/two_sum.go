package xsum

// TwoSumCountFast function counts pairs that sum to 0
func TwoSumCountFast(a []int) int {
	sorter.SortInts(a)

	if containsDuplicates(a) {
		panic("contains duplicate integers")
	}

	count := 0

	n := len(a)
	for i := 0; i < n; i++ {
		if j := binarySearch.Index(a, -a[i]); j > i {
			count++
		}
	}

	return count
}
