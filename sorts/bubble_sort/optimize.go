package bubble_sort

// OptimizedSort
// stable
// in-place
func OptimizedSort(sl []int) {
	for i := 0; i < len(sl)-1; i += 1 {
		isSwapped := false
		for j := 0; j < len(sl)-1-i; j += 1 {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
}
