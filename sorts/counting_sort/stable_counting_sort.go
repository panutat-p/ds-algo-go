package counting_sort

import "fmt"

/*
https://www.programiz.com/dsa/counting-sort
*/

// StableCountingSort
// not inplace
// stable
// harder to implement
func StableCountingSort(sl []int, min int, max int) []int {
	counts := make([]int, max-min+1)

	for _, v := range sl {
		index := v - min // this index point to count value
		counts[index] += 1
	}

	// use counts slice to store cumulative value
	for i := 1; i < len(counts); i += 1 {
		counts[i] += counts[i-1] // ⚠️ counts slice becomes cumulative slice
	}
	fmt.Printf("counts: %v\n\n", counts)

	// calculate index of result slice, then place the original value
	rslt := make([]int, len(sl))
	for i := len(sl) - 1; i >= 0; i -= 1 { // loop backward
		index := sl[i] - min          // this index point to cumulative value
		rslt[counts[index]-1] = sl[i] // 🦄 formula
		counts[index] -= 1            // decrease count until become 0, 🦊 stable sort
		fmt.Println("counts:", counts)
		fmt.Println("rslt:", rslt)
		fmt.Println()
	}

	return rslt
}
