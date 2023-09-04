package search

import (
	"fmt"
	"testing"
)

func TestFindSearchIterations(t *testing.T) {
	var table = []struct {
		lower    int
		upper    int
		item     int
		expected int
	}{
		{0, 16, 5, 3},
		{0, 100, 5, 5},
	}
	for _, x := range table {
		result := FindSearchIterations(x.lower, x.upper, x.item)
		if result != x.expected {
			t.Errorf("min:%d,max:%d,item:%d,result:%d,expected:%d\n", x.lower, x.upper, x.item, result, x.expected)
		}
	}
	fmt.Println("Done")
}
