// main is the functional test package
// Its primary objective is to calculate the times that the
// different algorithms take to sort the arrays
package main

import (
	"fmt"
	"math/rand"
	"time"

	"sorting/sorting"
	"sorting/sorting/bubble_sorting"
	"sorting/sorting/insertion_sorting"
	"sorting/sorting/selection_sorting"
	"sorting/sorting/standard_sorting"
)

type testingSorter struct {
	service     sorting.SortingService
	elapsedTime int64
	sorted      []int
	result      string
}

func main() {
	numbers := randomIntSlice(10000)
	sorters := createSorters()
	totalElapsedTime := int64(0)

	fmt.Printf("Array size: %d elements\n", len(numbers))

	for i := range sorters {
		sorter := &sorters[i]

		t := time.Now()
		sorter.sorted = sorter.service.Asc(numbers)
		sorter.elapsedTime = time.Since(t).Milliseconds()
		totalElapsedTime += sorter.elapsedTime

		sorter.result = "FAIL"
		if sorting.VerifyAscSortedArray(sorter.sorted) {
			sorter.result = "OK"
		}
	}

	for _, sorter := range sorters {
		incidence := float64(sorter.elapsedTime) / float64(totalElapsedTime) * 100

		fmt.Printf("[%s] - Duration: %dms (%.0f%% of total duration) - Result verification: %s\n",
			sorter.service.GetName(),
			sorter.elapsedTime,
			incidence,
			sorter.result)
	}
}

func createSorters() []testingSorter {
	return []testingSorter{
		{service: bubble_sorting.NewSorter()},
		{service: insertion_sorting.NewSorter()},
		{service: selection_sorting.NewSorter()},
		{service: standard_sorting.NewSorter()},
	}
}

func randomIntSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	ret := []int{}
	for i := 0; i < n; i++ {
		rnd := rand.Intn(1000000)
		ret = append(ret, rnd)
	}
	return ret
}
