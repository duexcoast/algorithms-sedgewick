package pq

import "sort"

type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants and may be called
// whenever the heap invariants have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		sink(h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h Interface, x any) {
	h.Push(x)
	swim(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n)where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h Interface) any {
	n := h.Len() - 1
	h.Swap(0, n)  // Swap the first and last elements
	sink(h, 0, n) // sink stops before n. So our swapped element is safe.
	return h.Pop()
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove(h Interface, i int) any {
	n := h.Len() - 1 // index of last element
	if n != i {      // check if i is the last element
		h.Swap(i, n) // swap i to end, to be popped after heapifying

		// since a heap only maintains partial order, we have to attempt to
		// both sink and swim the new item at i, we're not sure which way it
		// needs to go.
		if !sink(h, i, n) { // try to sink the new item at i
			swim(h, i) // if we can't sink it, try to let it swim
		}
	}
	// return the element
	return h.Pop()
}

// Fix re-establishes the heap orderign after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value
// The complexity is O(log n) where n = h.Len().
func Fix(h Interface, i int) {
	if !sink(h, i, h.Len()) {
		swim(h, i)
	}
}

func swim(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func sink(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1

		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2 // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		// follow the element down and try to sink it further
		i = j
	}
	return i > i0
}
