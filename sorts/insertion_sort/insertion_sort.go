package insertion_sort

func InsertionSort(sl []int) {
	for i := 1; i < len(sl); i += 1 {
		unsortedValue := sl[i]
		var j int
		for j = i - 1; j >= 0; j -= 1 { // loop through right to left of sorted part in the slice
			if sl[j] > unsortedValue {
				sl[j+1] = sl[j] // right shift
			} else { // found the stop condition
				break
			}
		}
		sl[j+1] = unsortedValue // insert behind the current j position, or insert to old position
	}
}

// RecursiveInsertionSort
// num: number of items
func RecursiveInsertionSort(sl []int, num int) {
	if num < 2 {
		return
	}

	RecursiveInsertionSort(sl, num-1)

	unsortedValue := sl[num-1]
	var i int
	for i = num - 2; i >= 0; i -= 1 {
		if sl[i] > unsortedValue {
			sl[i+1] = sl[i]
		} else {
			break
		}
	}
	sl[i+1] = unsortedValue
}
