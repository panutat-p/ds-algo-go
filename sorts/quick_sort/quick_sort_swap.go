package quick_sort

/*
https://www.programiz.com/dsa/quick-sort
*/

// QuickSortSwap
// inplace
// use value at lastIndex as pivot
// partitioning contains only one for loop
func QuickSortSwap(sl []int, firstIndex int, lastIndex int) {
	if firstIndex < lastIndex {
		pivotIndex := partitionSwap(sl, firstIndex, lastIndex)

		QuickSortSwap(sl, firstIndex, pivotIndex-1)
		QuickSortSwap(sl, pivotIndex+1, lastIndex)
	}
}

func partitionSwap(sl []int, firstIndex int, lastIndex int) int {
	pivot := sl[lastIndex]
	pivotIndex := firstIndex
	for i := firstIndex; i < lastIndex; i += 1 { // loop to element before the pivot
		if pivot >= sl[i] {
			sl[i], sl[pivotIndex] = sl[pivotIndex], sl[i] // lower values are swapped to front
			pivotIndex += 1                               // move forward
		}
	}
	sl[pivotIndex], sl[lastIndex] = sl[lastIndex], sl[pivotIndex] // swap pivot at lastIndex to pivotIndex
	return pivotIndex
}
