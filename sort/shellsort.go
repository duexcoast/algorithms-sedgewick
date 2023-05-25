package sort

func ShellSort(x Sortable) {
	n := x.Len()
	h := 1
	for h < n/3 {
		h = 3*h + 1 // 1, 4, 13, 40, 121, 364, 1093...
	}
	for h >= 1 {
		// h-sort the array
		for i := h; i < n; i++ {
			// insert a[i] among a[i-h], a[i-2*h], a[i-3*h]...
			for j := i; j >= h && x.Less(j, j-h); j -= h {
				x.Swap(j, j-h)
			}
		}
		h = h / 3
	}
}

type Shell struct{}

func NewShell() Shell {
	return Shell{}
}

// Implement Sorter
func (s Shell) SortInts(x []int) {
	ShellSort(IntSortSlice(x))
}
func (s Shell) SortFloat64s(x []float64) {
	ShellSort(Float64SortSlice(x))
}
func (s Shell) SortStrings(x []string) {
	ShellSort(StringSortSlice(x))
}
