package uf_test

import (
	"fmt"
	"testing"

	"github.com/duexcoast/algorithms-sedgewick/fundamentals/uf"
	"github.com/duexcoast/algorithms-sedgewick/testutil"
)

func TestQuickFind(t *testing.T) {
	in := testutil.NewInReadWords("testdata/tinyUF.txt")

	n := in.ReadInt()
	unionFind := uf.NewQuickFindUF(n)

	for !in.IsEmpty() {
		p := in.ReadInt()
		q := in.ReadInt()

		if unionFind.Connected(p, q) {
			continue
		}
		unionFind.Union(p, q)

		fmt.Printf("%d %d", p, q)
	}

	fmt.Printf("%d components", unionFind.Count())
}
