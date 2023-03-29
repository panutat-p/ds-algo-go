package binary_search

import "fmt"

// BinarySearchRecursive
// s >> [11]
// s[1:] >> []
// s[:0] >> []
func BinarySearchRecursive(li []int, target int) bool {
	if len(li) == 0 { // base case for empty list and not found
		return false
	}

	midIndex := len(li) / 2
	if li[midIndex] == target {
		return true
	} else if li[midIndex] < target {
		fmt.Println(li[midIndex+1:])
		return BinarySearchRecursive(li[midIndex+1:], target)
	} else {
		fmt.Println(li[:midIndex])
		return BinarySearchRecursive(li[:midIndex], target)
	}
}
