package selection_sort

import (
	"fmt"
	"time"
)

func SelectionSort(sl []int) {
	start := time.Now()
	count := 0

	for i := len(sl) - 1; i > 0; i -= 1 {
		index := 0
		for j := 1; j <= i; j += 1 {
			if sl[j] > sl[index] {
				index = j
			}
			count += 1
		}
		sl[index], sl[i] = sl[i], sl[index]
	}

	fmt.Println("count:", count)
	elapsed := time.Since(start)
	fmt.Println("time elapsed:", elapsed)
}
