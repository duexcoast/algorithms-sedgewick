package uf_test

import (
	"testing"

	"github.com/duexcoast/algorithms-sedgewick/fundamentals/uf"
	"github.com/duexcoast/algorithms-sedgewick/testutil"
)

// go test -v -run="none" -bench="." -benchtime="3s"

type pair struct {
	p, q int
}

type testCase struct {
	n     int
	pairs []pair
}

var testdata *testCase

func dataInit() {
	in := testutil.NewInReadWords("testdata/mediumUF.txt")

	n := in.ReadInt()
	pairs := make([]pair, 1)

	for !in.IsEmpty() {
		p := in.ReadInt()
		q := in.ReadInt()
		pairs = append(pairs, pair{p, q})
	}

	testdata = &testCase{n, pairs}

}

func benchUF(unionFind uf.UnionFind, pairs []pair) {
	for _, pq := range pairs {
		if unionFind.Connected(pq.p, pq.q) {
			continue
		}
		unionFind.Union(pq.p, pq.q)
	}
}
func BenchmarkQuickUnionUF(b *testing.B) {
	b.ReportAllocs()
	if testdata == nil {
		b.StopTimer()
		dataInit()
		b.StartTimer()
	}

	unionFind := uf.NewQuickUnionUF(testdata.n)

	for i := 0; i < b.N; i++ {
		benchUF(unionFind, testdata.pairs)
	}
}

func BenchmarkQuickFindUF(b *testing.B) {
	b.ReportAllocs()
	if testdata == nil {
		b.StopTimer()
		dataInit()
		b.StartTimer()
	}

	unionFind := uf.NewQuickFindUF(testdata.n)

	for i := 0; i < b.N; i++ {
		benchUF(unionFind, testdata.pairs)
	}
}
