// QuickUnionUF Quick-union algorithm
package uf

type QuickUnionUF struct {
	id    []int // id[i]: component identifier of i
	count int   // number of components
}

func NewQuickUnionUF(n int) *QuickUnionUF {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return &QuickUnionUF{id: id, count: n}
}

func (qf *QuickUnionUF) Find(p int) int {
	// Find component name by following links until we reach
	// reach a site that has a link to itself - this is the root.
	for p != qf.id[p] {
		p = qf.id[p]
	}
	return p
}

func (qf *QuickUnionUF) Union(p, q int) {
	// Give p and q the same root by linking one root to the other.
	// We arbitrarily decide to link p to q, by renaming p's root to q's root
	pRoot := qf.Find(p)
	qRoot := qf.Find(q)

	if pRoot == qRoot {
		return
	}

	qf.id[pRoot] = qRoot

	qf.count--
}

func (qf *QuickUnionUF) Count() int {
	return qf.count
}

func (qf *QuickUnionUF) Connected(p, q int) bool {
	return qf.Find(p) == qf.Find(q)
}
