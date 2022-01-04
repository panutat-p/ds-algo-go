package quick_sort

import (
	"fmt"
	"testing"
)

func IsSorted(sl []int) bool {
	for i := 0; i < len(sl)-1; i += 1 {
		if sl[i] > sl[i+1] {
			return false
		}
	}
	return true
}

func TestQuickSortCross(t *testing.T) {
	sl := []int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63}
	QuickSortCross(sl, 0, len(sl)) // ⚠️ pass len index
	fmt.Println(sl)

	if !IsSorted(sl) {
		t.Errorf("got %v, but expect %v", IsSorted(sl), true)
	}
}

func TestQuickSortCross2(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	QuickSortCross(sl, 0, len(sl)) // ⚠️ pass len index
	fmt.Println(sl)

	if !IsSorted(sl) {
		t.Errorf("got %v, but expect %v", IsSorted(sl), true)
	}
}
