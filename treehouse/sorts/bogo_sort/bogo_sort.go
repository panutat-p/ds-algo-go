package bogo_sort

import (
	"fmt"
	"math/rand"
	"time"
)

func IsSorted(sl []int) bool {
	for i := 0; i < len(sl)-1; i += 1 {
		if sl[i] > sl[i+1] {
			return false
		}
	}
	return true
}

func BogoSort(sl []int) []int {
	start := time.Now()
	count := 0
	for !IsSorted(sl) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(sl), func(i int, j int) {
			sl[i], sl[j] = sl[j], sl[i]
		})
		count += 1
	}
	fmt.Println("count:", count)
	elapsed := time.Since(start)
	fmt.Println("time elapsed:", elapsed)
	return sl
}
