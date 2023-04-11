package main

import "fmt"

func main() {
	fmt.Println("🟩 want [1 2 4 5]", "got", quickSort([]int{2, 5, 1, 4}))
}

func quickSort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	if len(nums) == 1 {
		return nums
	}

	pivot := nums[0]
	var (
		left  []int
		right []int
	)
	for _, v := range nums[1:] {
		if pivot > v {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}
