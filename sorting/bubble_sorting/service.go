package bubble_sorting

type sorter struct {
}

func NewSorter() sorter {
	return sorter{}
}

func (sorter) GetName() string {
	return "Bubble Sort"
}

func (sorter) Asc(unsortedItems []int) []int {
	// Copy the original array into a new one. This allows to preserve the unsorted values
	items := make([]int, len(unsortedItems))
	copy(items, unsortedItems)

	for i := 0; i < len(items); i++ {
		swapped := false
		for j := len(items) - 1; j > i; j-- {
			if items[j] < items[j-1] {
				items[j], items[j-1] = items[j-1], items[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return items
}
