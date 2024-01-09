package bubble_sort

// BubbleUp
// stable
// in-place
func BubbleUp(nums []int) []int {
	for i := 0; i < len(nums); i += 1 {
		for j := 0; j < len(nums)-1-i; j += 1 {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
