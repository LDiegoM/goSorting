package standard_sorting

type ascSorter []int

func (a ascSorter) Len() int           { return len(a) }
func (a ascSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ascSorter) Less(i, j int) bool { return a[i] < a[j] }
