package main

import (
	"github.com/duexcoast/algorithms-sedgewick/fundamentals/xsum"
	"github.com/duexcoast/algorithms-sedgewick/testutil"
)

var (
	rand testutil.Random
	// ratio = flag.Bool("r", false, "doubling ratio test")
)

func init() {
	rand = *testutil.NewRandom()
}

func main() {
	// flag.Parse()

	DoublingRatio()
	// doublingTest()
}

const maxInteger = 100000

func timeTrial(n int) float64 {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.UniformIntRange(-maxInteger, maxInteger)
	}

	timer := testutil.NewStopwatch()

	xsum.ThreeSumCount(a)

	return timer.ElapsedTime()
}
