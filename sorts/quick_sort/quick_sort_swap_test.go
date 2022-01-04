package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSortSwap(t *testing.T) {
	sl := []int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63}
	QuickSortSwap(sl, 0, len(sl)-1) // ⚠️ pass last index
	fmt.Println(sl)

	if !IsSorted(sl) {
		t.Errorf("got %v, but expect %v", IsSorted(sl), true)
	}
}

func TestQuickSortSwap2(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	QuickSortSwap(sl, 0, len(sl)-1) // ⚠️ pass last index
	fmt.Println(sl)

	if !IsSorted(sl) {
		t.Errorf("got %v, but expect %v", IsSorted(sl), true)
	}
}
