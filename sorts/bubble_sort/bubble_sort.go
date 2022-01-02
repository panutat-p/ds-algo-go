package bubble_sort

import (
	"fmt"
	"time"
)

func BubbleSort(sl []int) {
	start := time.Now()
	count := 0

	for i := 0; i < len(sl)-1; i += 1 {
		for j := 0; j < len(sl)-1; j += 1 {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
			count += 1
		}
	}

	fmt.Println("count:", count)
	elapsed := time.Since(start)
	fmt.Println("time elapsed:", elapsed)
}

func OptimizedBubbleSort(sl []int) {
	start := time.Now()
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
	elapsed := time.Since(start)
	fmt.Println("time elapsed:", elapsed)
}

func IsSorted(sl []int) bool {
	for i := 0; i < len(sl)-1; i += 1 {
		if sl[i] > sl[i+1] {
			return false
		}
	}
	return true
}
