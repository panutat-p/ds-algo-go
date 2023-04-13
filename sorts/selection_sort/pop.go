package selection_sort

import (
	"fmt"
	"math"
	"time"
)

// Sort3 O(n^2) time
// compare to max int
// remove the targeted element from the slice until its length become zero
func Sort3(sl []int) []int {
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
	fmt.Println("time elapsed:", time.Since(start))
	return sortedSl
}
