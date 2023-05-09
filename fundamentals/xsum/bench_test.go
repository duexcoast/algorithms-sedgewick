package xsum_test

import (
	"testing"

	. "github.com/duexcoast/algorithms-sedgewick/fundamentals/xsum"
	"github.com/duexcoast/algorithms-sedgewick/testutil"
)

// % go test -v -run="none" -bench="." -benchtime="3s"

var a = testutil.NewIn("testdata/1Kints.txt").ReadAllInts()

func BenchmarkThreeSumCount(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ThreeSumCount(a[0:])
	}
}

func BenchmarkThreeSumCountFast(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ThreeSumCountFast(a[0:])
	}
}
