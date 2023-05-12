// WeightedQuickUnionUF Weighted quick-union algorithm
package uf

type WeightedQuickUnionUF struct {
	id    []int // id[i]: component identifier of i
	size  []int // size of component for roots (site indexed)
	count int   // number of components
}

func NewWeightedQuickUnionUF(n int) *WeightedQuickUnionUF {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sz := make([]int, n)
	for i := 0; i < n; i++ {
		sz[i] = 1
	}

	return &WeightedQuickUnionUF{id: id, size: sz, count: n}
}

func (qf *WeightedQuickUnionUF) Find(p int) int {
	// Find component name by following links until we reach
	// reach a site that has a link to itself - this is the root.
	for p != qf.id[p] {
		p = qf.id[p]
	}
	return p
}

func (qf *WeightedQuickUnionUF) Union(p, q int) {
	// Give p and q the same root by linking one root to the other.
	// We arbitrarily decide to link p to q, by renaming p's root to q's root
	i := qf.Find(p)
	j := qf.Find(q)

	if i == j {
		return
	}

	// make smaller root point to larger one
	if qf.size[i] < qf.size[j] {
		qf.id[i] = qf.id[j]
		qf.size[j] += qf.size[i]
	} else {
		qf.id[j] = qf.id[i]
		qf.size[i] += qf.id[j]
	}

	qf.count--
}

func (qf *WeightedQuickUnionUF) Count() int {
	return qf.count
}

func (qf *WeightedQuickUnionUF) Connected(p, q int) bool {
	return qf.Find(p) == qf.Find(q)
}
