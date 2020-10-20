package math

// Sortable represents objects that can be sorted.
type Sortable interface {
	// Len returns the number of elements in the collection.
	Size() int
	// Less reports whether the element with index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Sort sorts data by quick sort.
func Sort(data Sortable) {
	QuickSort(data)
}

// QuickSort sorts data by quick sort.
func QuickSort(data Sortable) {
	QuickSortRange(data, 0, data.Size() - 1)
}

// QuickSortRange sorts data[l, r] by quick sort.
func QuickSortRange(data Sortable, l int, r int) {
	if l < r {
		index := partition(data, l, r)
		QuickSortRange(data, l, index - 1)
		QuickSortRange(data, index + 1, r)
	}
}

func partition(data Sortable, l int, r int) int {
	t := r
	for l < r {
		for l < r && data.Less(l, t) {
			l++
		}
		for l < r && !data.Less(r, t) {
			r--
		}
		data.Swap(l, r)
	}
	data.Swap(l, t)
	return l
}