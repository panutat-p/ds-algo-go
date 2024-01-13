package main

import "fmt"

var count int

func main() {
	sl := []int{7, 5, 4, 13, 9, 10, 54, -4, 3, 8, 6, 1, 2}
	sl = sort(sl)
	fmt.Println(count)
	fmt.Println(sl)
}

func sort(nums []int) []int {
	// edge
	if len(nums) < 2 {
		return nums
	}

	// pivot
	var (
		left  []int
		right []int
	)
	pivot := nums[0]
	for _, n := range nums[1:] {
		if pivot > n {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
		count += 1
	}

	// repeat
	left = sort(left)
	right = sort(right)

	// join
	var ret []int
	ret = append(ret, left...)
	ret = append(ret, pivot)
	ret = append(ret, right...)
	return nums
}

// bubble 156
// bubble optimized 78
// selection 78
// quick 34
// merge ?
