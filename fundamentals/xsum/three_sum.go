// Package xsum provides functions counting the number of triples in an array of
// N integers that sum to 0
package xsum

func ThreeSumCount(a []int) int {
	n := len(a)
	count := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < j; k++ {
				if a[i]+a[j]+a[k] == 0 {
					count++
				}
			}
		}
	}

	return count
}
