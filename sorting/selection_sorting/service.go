package selection_sorting

type sorter struct {
}

func NewSorter() sorter {
	return sorter{}
}

func (sorter) GetName() string {
	return "Selection Sort"
}

func (sorter) Asc(unsortedItems []int) []int {
	// Copy the original array into a new one. This allows to preserve the unsorted values
	items := make([]int, len(unsortedItems))
	copy(items, unsortedItems)

	for i := 0; i < len(items); i++ {
		k := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[k] {
				k = j
			}
		}

		items[i], items[k] = items[k], items[i]
	}

	return items
}
