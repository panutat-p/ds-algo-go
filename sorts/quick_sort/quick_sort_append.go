package quick_sort

import "fmt"

// QuickSortAppend
// use additional space: left []int, right []int
// easy to implement
func QuickSortAppend(sl []int) []int {
	if len(sl) <= 1 {
		return sl
	}

	var left []int
	var right []int
	pivot := sl[0]
	for _, v := range sl[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	fmt.Printf("%v  %v  %v\n", left, pivot, right)

	left = QuickSortAppend(left)
	right = QuickSortAppend(right)
	return append(append(left, pivot), right...)
}
