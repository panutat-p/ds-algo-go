package main

import (
	"fmt"
	"slices"
)

var count int

func main() {
	sl := []int{7, 5, 4, 13, 9, 10, 54, -4, 3, 8, 6, 1, 2}
	sl = sort(sl)
	fmt.Println(count)
	fmt.Println(sl)
	if !slices.IsSorted(sl) {
		panic("not sorted")
	}
}

func sort(nums []int) []int {
	// edge
	if len(nums) < 2 {
		return nums
	}

	// split
	left := nums[:len(nums)/2]
	right := nums[len(nums)/2:]

	// repeat
	left = sort(left)
	right = sort(right)

	var (
		i   int
		j   int
		ret []int
	)
	// merge
	for i < len(left) && j < len(right) {
		count += 1
		if left[i] <= right[j] {
			ret = append(ret, left[i])
			i += 1
		} else {
			ret = append(ret, right[i])
			j += 1
		}
	}

	// merge remaining
	ret = append(ret, left[i:]...)
	ret = append(ret, right[j:]...)
	return ret
}

// bubble 156 comparisons
// bubble optimized 78 comparisons
// selection 78 comparisons
// quick 34 comparisons, 17 call stacks
// merge 34 comparisons, 25 call stacks
