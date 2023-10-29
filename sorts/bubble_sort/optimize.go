package bubble_sort

import (
	"fmt"
)

func OptimizedSort(sl []int) {
	count := 0

	for i := 0; i < len(sl)-1; i += 1 {
		isSwapped := false
		for j := 0; j < len(sl)-1-i; j += 1 {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
				isSwapped = true
			}
			count += 1
		}
		if !isSwapped { // already sorted
			break
		}
	}

	fmt.Println("count:", count)
}
