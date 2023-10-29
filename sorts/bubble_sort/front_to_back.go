package bubble_sort

// FrontToBack
// stable
// in-place
func FrontToBack(nums []int) []int {
	for i := 0; i < len(nums); i += 1 {
		for j := i + 1; j < len(nums); j += 1 {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
