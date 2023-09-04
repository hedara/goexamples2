package sorting

import (
	"fmt"
	"testing"
)

func TestQuicksort(t *testing.T) {
	var table = []struct {
		array  []int
		sorted []int
	}{
		{
			[]int{8, 4, 6, 1, 7, 3, 2, 5},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{20, 43, 16, 10, 0, 8},
			[]int{0, 8, 10, 16, 20, 43},
		},
		{
			[]int{9, 3, 18, 19, 19, 4, 1},
			[]int{1, 3, 4, 9, 18, 19, 19},
		},
	}
	for _, x := range table {
		y_array := make([]int, len(x.array))
		copy(y_array, x.array)
		Quicksort(&y_array, 0, len(y_array)-1)
		if !compareSlices(&y_array, &x.sorted) {
			t.Errorf("Failure sorting values\n")
		}
		fmt.Println()
	}
}

func compareSlices(a, b *[]int) bool {
	if len((*a)) != len((*b)) {
		return false
	}
	for i, v := range *a {
		if (*b)[i] != v {
			return false
		}
	}
	return true
}
