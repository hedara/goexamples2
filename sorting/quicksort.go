package sorting

import (
	"math/rand"
	"time"
)

var rseed = rand.NewSource(time.Now().UnixNano())
var rgenerator = rand.New(rseed)

func Quicksort(nums *[]int, start int, end int) {
	if start >= end {
		return
	}
	pivotIndex := start + rgenerator.Intn(end-start+1)
	pivotIndex = partition(nums, pivotIndex, start, end)
	Quicksort(nums, 0, pivotIndex-1)
	Quicksort(nums, pivotIndex+1, end)
}

func partition(nums *[]int, pivotIndex int, start int, end int) int {
	pIndex := start
	pivotVal := (*nums)[pivotIndex]
	for i := start; i <= end; i++ {
		if (*nums)[i] < pivotVal {
			if i != pIndex {
				(*nums)[i], (*nums)[pIndex] = (*nums)[pIndex], (*nums)[i]
			}
			pIndex += 1
		}
	}
	if pIndex < pivotIndex {
		(*nums)[pIndex], (*nums)[pivotIndex] = (*nums)[pivotIndex], (*nums)[pIndex]
	}
	return pIndex
}
