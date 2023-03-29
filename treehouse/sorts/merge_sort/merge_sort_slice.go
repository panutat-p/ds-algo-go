package merge_sort

import "fmt"

// MergeSort O(kn log n) time, O(n) space
// copy original slice
func MergeSort(sl []int) []int {
	if len(sl) <= 1 {
		return sl
	}

	left, right := Split(sl)
	fmt.Println("split slices:", left, right)
	left = MergeSort(left)
	right = MergeSort(right)
	return Merge(left, right)
}

// Split O(k log n) time
// create 2 new slices from the original
func Split(sl []int) ([]int, []int) {
	return sl[:len(sl)/2], sl[len(sl)/2:]
}

// Merge O(n) time
// concatenate 2 sorted slices by picking lower value first
func Merge(left []int, right []int) []int {
	var sl []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sl = append(sl, left[i])
			i += 1
		} else {
			sl = append(sl, right[j])
			j += 1
		}
	}

	// if left or right are not equal size
	// append remaining elements
	sl = append(sl, left[i:]...)
	sl = append(sl, right[j:]...)

	fmt.Println("merged:", sl)
	return sl
}

func IsSorted(sl []int) bool {
	length := len(sl)
	if length == 0 || length == 1 {
		return true
	}

	return sl[0] <= sl[1] && IsSorted(sl[1:])
}
