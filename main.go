package main

import (
	"fmt"
	"job2/search"
	"job2/sorting"
)

func main() {
	nums := []int{8, 4, 6, 1, 7, 3, 2, 5}

	sorting.Quicksort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
	// binary search
	// runBinarySearch()

}

func runBinarySearch() {
	result := search.FindSearchIterations(0, 10000, 5)
	fmt.Println(result)

}
