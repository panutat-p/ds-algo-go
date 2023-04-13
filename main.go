package main

import "fmt"

func main() {
	fmt.Println("🟩 want [1 2 4 5]", "got", selectionSort([]int{2, 5, 1, 4}))
}

func selectionSort(nums []int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j += 1 {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				fmt.Println(nums)
			}
		}
	}
	return nums
}
