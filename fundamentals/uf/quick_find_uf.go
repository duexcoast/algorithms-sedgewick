package uf

type QuickFindUF struct {
	id    []int // id[i]: component identifier of i
	count int   // number of components
}

func NewQuickFindUF(n int) *QuickFindUF {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return &QuickFindUF{id: id, count: n}
}

func (qf *QuickFindUF) Find(p int) int {
	return qf.id[p]
}

func (qf *QuickFindUF) Union(p, q int) {
	// Put p and q into the same component
	pID := qf.Find(p)
	qID := qf.Find(q)

	// if p and q are already in the same component then nothing to do
	if pID == qID {
		return
	}

	// Rename p's component to q's name
	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pID {
			qf.id[i] = qID
		}
	}
	qf.count--

}

func (qf *QuickFindUF) Count() int {
	return qf.count
}

func (qf *QuickFindUF) Connected(p, q int) bool {
	return qf.Find(p) == qf.Find(q)
}

// func (qf *QuickFindUF) validate(p int) {}
