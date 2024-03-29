package bubble_sort

// PlainBubbleSort
// stable
// in-place
func PlainBubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i += 1 {
		for j := 0; j < len(nums)-1; j += 1 {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
