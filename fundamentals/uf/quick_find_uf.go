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

	return 0
}

func (qf *QuickFindUF) Union(p, q int) {

}

func (qf *QuickFindUF) Count() int {

	return 0
}

func (qf *QuickFindUF) Connected(p, q int) bool {
	return find(p) == find(q)
}

func (qf *QuickFindUF) validate(p int) {}
