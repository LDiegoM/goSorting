package standard_sorting

import "sort"

type sorter struct {
}

func NewSorter() sorter {
	return sorter{}
}

func (sorter) GetName() string {
	return "Standard lib Sort"
}

func (sorter) Asc(unsortedItems []int) []int {
	// Copy the original array into a new one. This allows to preserve the unsorted values
	items := make([]int, len(unsortedItems))
	copy(items, unsortedItems)

	sort.Sort(ascSorter(items))

	return items
}
