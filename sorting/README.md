# Sorting algorithms implementations

This package implements different methods of sorting arrays in go.

Read this [page](https://www.toptal.com/developers/sorting-algorithms) to get information about the different algorithms implemented.

The different methods must implement the following interface:

``` go
type SortingService interface {
	Asc(unsortedItems []int) []int
	GetName() string
}
```

## Bubble sorting

Bubble sort has many of the same properties as insertion sort, but has slightly higher overhead. 

In the case of nearly sorted data, bubble sort takes O(n) time, but requires at least 2 passes through the data (whereas insertion sort requires something more like 1 pass).

## Insertion sorting

Although it is one of the elementary sorting algorithms with O(n2) worst-case time, insertion sort is the algorithm of choice either when the data is nearly sorted (because it is adaptive) or when the problem size is small (because it has low overhead).

For these reasons, and because it is also stable, insertion sort is often used as the recursive base case (when the problem size is small) for higher overhead divide-and-conquer sorting algorithms, such as merge sort or quick sort.

## Selection sorting

From the comparions presented here, one might conclude that selection sort should never be used. It does not adapt to the data in any way (notice that the four animations above run in lock step), so its runtime is always quadratic.

However, selection sort has the property of minimizing the number of swaps. In applications where the cost of swapping items is high, selection sort very well may be the algorithm of choice.
