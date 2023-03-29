package binary_search

import "fmt"

// BinarySearch
// li must be sorted
func BinarySearch(li []int, target int) int {
	firstIndex := 0
	lastIndex := len(li) - 1

	for firstIndex <= lastIndex {
		midIndex := (firstIndex + lastIndex) / 2
		fmt.Println("midIndex:", midIndex)

		if li[midIndex] == target {
			return midIndex
		} else if li[midIndex] < target {
			firstIndex = midIndex + 1
		} else {
			lastIndex = midIndex - 1
		}
	}
	return -1
}
