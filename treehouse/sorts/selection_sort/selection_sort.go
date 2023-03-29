package selection_sort

import (
	"fmt"
	"math"
	"time"
)

// SelectionSort O(n^2) time
func SelectionSort(sl []int) []int {
	start := time.Now()
	var sortedSl []int
	for len(sl) != 0 {
		var index int
		min := math.MaxInt
		for i, v := range sl {
			if min > v {
				index = i
				min = v
			}
		}
		sl = append(sl[:index], sl[(index+1):]...)
		sortedSl = append(sortedSl, min)
	}
	elapsed := time.Since(start)
	fmt.Println("time elapsed:", elapsed)
	return sortedSl
}
