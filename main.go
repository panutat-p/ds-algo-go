package main

import "fmt"

func main() {
	fmt.Println("🟩 want [1 2 4 5]", "got", quickSort([]int{2, 5, 1, 4}))
}

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i += 1 {
		for j := i + 1; j < len(nums); j += 1 {
			if nums[i] > nums[j] {
				// found lower value in the right
				nums[i], nums[j] = nums[j], nums[i]
				fmt.Println("nums", nums)
			}
		}
	}
	return nums
}

func quickSort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var (
		left  []int
		right []int
	)
	pivot := nums[0]
	for _, v := range nums[1:] {
		if pivot > v {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)
	result := append(append(left, pivot), right...)
	return result
}
