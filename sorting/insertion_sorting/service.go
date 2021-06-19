package insertion_sorting

type sorter struct {
}

func NewSorter() sorter {
	return sorter{}
}

func (sorter) GetName() string {
	return "Insertion Sort"
}

func (sorter) Asc(unsortedItems []int) []int {
	// Copy the original array into a new one. This allows to preserve the unsorted values
	items := make([]int, len(unsortedItems))
	copy(items, unsortedItems)

	for i := 1; i < len(items); i++ {
		for k := i; k > 0 && items[k] < items[k-1]; k-- {
			items[k], items[k-1] = items[k-1], items[k]
		}
	}

	return items
}
