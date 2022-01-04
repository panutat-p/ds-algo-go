package quick_sort

import "fmt"

// QuickSortCross
// inplace
// use value at startIndex as pivot
// partitioning iterate both i and j from both start and end
func QuickSortCross(sl []int, startIndex int, lenIndex int) {
	if lenIndex-startIndex < 2 { // empty slice or one element slice
		return
	}

	pivotIndex := partitionCross(sl, startIndex, lenIndex)
	fmt.Printf("%v   %v   %v\n", sl[:pivotIndex], sl[pivotIndex], sl[pivotIndex+1:])

	QuickSortCross(sl, startIndex, pivotIndex)
	QuickSortCross(sl, pivotIndex+1, lenIndex)
}

func partitionCross(sl []int, startIndex int, lenIndex int) int {
	// at the beginning, j = len(sl) is index out of bound
	pivot := sl[startIndex] // use first element as a pivot
	i := startIndex
	j := lenIndex
	for i < j { // cross when i = j then end
		// do j first because we use i = 0 as pivot index
		for i < j {
			j -= 1 // move j backward
			if sl[j] < pivot {
				// j does not cross i yet
				sl[i] = sl[j] // assign right value to the left of pivot
				break
			}
		}

		for i < j {
			i += 1 // move i forward, after this line, start at i = 1
			if sl[i] > pivot {
				// i does not cross j yet
				sl[j] = sl[i] // assign left value to the right of pivot
				break
			}
		}
	}
	sl[j] = pivot // when i cross j, place the pivot here and return this index
	return j
}
