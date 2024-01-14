package merge_sort

// MergeSortSlice O(kn log n) time, O(n) space
// copy original slice
func MergeSortSlice(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// split
	// O(k log n) time
	left := nums[:len(nums)/2]
	right := nums[len(nums)/2:]

	// repeat
	left = MergeSortSlice(left)
	right = MergeSortSlice(right)

	// merge
	// O(n) time
	var (
		i   int
		j   int
		ret []int
	)
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			ret = append(ret, left[i])
			i += 1
		} else {
			ret = append(ret, right[j])
			j += 1
		}
	}

	// merge remaining
	ret = append(ret, left[i:]...)
	ret = append(ret, right[j:]...)

	return ret
}
