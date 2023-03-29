package quick_sort

import "fmt"

// QuickSort
// worst case O(n^2): pivot is always the biggest value
// best case O(n log n): pivot is always the midpoint value
// improve by using random pivot
func QuickSort(sl []int) []int {
	if len(sl) <= 1 { // base case
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
	fmt.Printf("%v  <%v<  %v\n", left, pivot, right)
	tmp := append(QuickSort(left), pivot)
	return append(tmp, QuickSort(right)...)
}
