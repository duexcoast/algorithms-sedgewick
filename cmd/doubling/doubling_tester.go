package main

import "fmt"

// doublingTest function generates a sequence of random input arrays, doubling
// the array size at each step, and prints out the running times of ThreeSumCount
// for each input size.
func doublingTest() {
	for n := 250; true; n += n {
		time := timeTrial(n)
		fmt.Printf("%7d %7.1f\n", n, time)
	}
}
