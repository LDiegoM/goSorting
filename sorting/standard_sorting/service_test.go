package standard_sorting

import (
	"testing"

	"sorting/sorting"

	"github.com/stretchr/testify/assert"
)

func TestAsc(t *testing.T) {
	testCases := sorting.TestCases

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			service := NewSorter()

			originalSlice := make([]int, len(tc.InputSlice))
			copy(originalSlice, tc.InputSlice)

			output := service.Asc(originalSlice)
			assert.Equal(t, tc.ExpectedOutput, output)

			// Verifies that the algorithm made a copy of the input slice
			assert.Equal(t, tc.InputSlice, originalSlice)
		})
	}
}
