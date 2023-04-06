package main

import "fmt"

func main() {
	fmt.Println("🟩 want [-1,2,3]", "got", selectionSort([]int{3, 2, -1}))
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
