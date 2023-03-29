package bogo_sort

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/panutat-p/fiset-complete-ds-go/pkg"
)

func BogoSort(sl []int) []int {
	start := time.Now()
	count := 0
	for !pkg.IsSorted(sl) {
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
