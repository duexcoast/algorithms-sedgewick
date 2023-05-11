package uf_test

import (
	"fmt"

	"github.com/duexcoast/algorithms-sedgewick/fundamentals/uf"
	"github.com/duexcoast/algorithms-sedgewick/testutil"
)

func ExampleQuickFindUF() {
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
		fmt.Printf("%d %d\n", p, q)
	}

	fmt.Printf("%d components\n", unionFind.Count())

	// Output:
	// 4 3
	// 3 8
	// 6 5
	// 9 4
	// 2 1
	// 5 0
	// 7 2
	// 6 1
	// 2 components

}
