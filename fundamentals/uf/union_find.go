package uf

type UnionFind interface {
	// Union adds a connection between p and q
	Union(p, q int)

	// Find returns component identifier for p (0 to N-1)
	Find(p int) int

	// Connected returns true if p and q are in the same component
	Connected(p, q int) bool

	// Count returns the number of components
	Count() int
}
