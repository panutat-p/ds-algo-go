package main

import "fmt"

func main() {
	sl := []int{5, 4, 9, 6, 1, 2}
	sl = quickSort(sl)
	fmt.Println(sl)
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	var (
		left  []int
		right []int
		pivot = nums[0]
	)
	for _, e := range nums[1:] {
		if pivot > e {
			left = append(left, e)
		} else {
			right = append(right, e)
		}
	}

	var (
		ret []int
	)
	ret = append(quickSort(left), pivot)
	ret = append(ret, quickSort(right)...)
	return ret
}
