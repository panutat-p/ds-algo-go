package main

import "fmt"

func main() {
	sl := []int{5, 4, 9, 6, 1, 2}
	selectionSort(sl)
	fmt.Println(sl)
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
