package main

import "fmt"

func main() {
	nums := []int{5, 3, 1, 9, 4, 8, 6, 2, 7}
	selectionSort(nums)
	fmt.Println("🟩", nums)
}

func selectionSort(nums []int) {
	for i := 0; i < len(nums); i += 1 {
		for j := i + 1; j < len(nums); j += 1 {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
