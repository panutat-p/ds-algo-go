package selection_sort

// Sort1
// pin the first element
func Sort1(nums []int) []int {
	for i := 0; i < len(nums); i += 1 {
		for j := i + 1; j < len(nums); j += 1 {
			if nums[i] > nums[j] {
				// found lower value in the right
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// Sort2
// pin the first element
// loop backward
func Sort2(sl []int) {
	for i := len(sl) - 1; i > 0; i -= 1 {
		index := 0
		for j := 1; j <= i; j += 1 {
			if sl[j] > sl[index] {
				index = j
			}
		}
		sl[index], sl[i] = sl[i], sl[index]
	}
}
