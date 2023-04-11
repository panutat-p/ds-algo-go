package main

import "fmt"

func main() {
	nums := []int{2, 5, 1, 4}
	selectionSort(nums)
	fmt.Println("🟩 want [1 2 4 5]", "got", nums)
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
