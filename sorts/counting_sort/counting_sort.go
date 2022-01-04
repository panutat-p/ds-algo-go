package counting_sort

// CountingSort
// not inplace
// need to know min value and max value in slice
// only valid for zero and positive numbers
func CountingSort(sl []int, min int, max int) {
	counts := make([]int, max-min+1)
	for i := 0; i < len(sl); i += 1 {
		index := sl[i] - min // for s[i]=1, min=1, 1-1=0, count 1 at index 0
		counts[index] += 1
	}

	i := 0
	for value := min; value <= max; value += 1 {
		index := value - min
		for counts[index] > 0 {
			sl[i] = value // place sorted value
			i += 1        // right shift on the result

			counts[index] -= 1 // reduce count until 0
		}
	}
}
