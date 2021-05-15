package sorting

type TestCase struct {
	Name           string
	InputSlice     []int
	ExpectedOutput []int
}

var TestCases = []TestCase{
	{
		Name:           "Given_Many_Items_Array_Then_Return_It_Sorted",
		InputSlice:     []int{1, 2, 5, 30, 19, 55, 2, 4, 3, 32, 28, 78, 59, 11, 18},
		ExpectedOutput: []int{1, 2, 2, 3, 4, 5, 11, 18, 19, 28, 30, 32, 55, 59, 78},
	},
	{
		Name:           "Given_One_Item_Array_Then_Return_The_Same_Array",
		InputSlice:     []int{12},
		ExpectedOutput: []int{12},
	},
	{
		Name:           "Given_Empty_Array_Then_Return_Empty_Array",
		InputSlice:     []int{},
		ExpectedOutput: []int{},
	},
}

type SortingService interface {
	Asc(unsortedItems []int) []int
	GetName() string
}

func VerifyAscSortedArray(items []int) bool {
	size := len(items)
	for i := 1; i < size; i++ {
		if items[i-1] > items[i] {
			return false
		}
	}
	return true
}
