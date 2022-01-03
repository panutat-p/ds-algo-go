package shell_sort

/*
https://www.programiz.com/dsa/shell-sort
*/

func ShellSort(sl []int) {
	for gap := len(sl) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(sl); i += 1 {
			unsortedValue := sl[i]
			var j int
			for j = i; j >= gap; j -= gap {
				if sl[j-gap] > unsortedValue {
					sl[j] = sl[j-gap] // right shift with 3 positions
				} else {
					break
				}
			}
			sl[j] = unsortedValue
		}
	}
}
