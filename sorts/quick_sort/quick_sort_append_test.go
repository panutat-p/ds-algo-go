package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSortAppend(t *testing.T) {
	sl := []int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63}
	rslt := QuickSortAppend(sl)
	fmt.Println(rslt)

	if !IsSorted(rslt) {
		t.Errorf("got %v, but expect %v", IsSorted(rslt), true)
	}
}

func TestQuickSortAppend2(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rslt := QuickSortAppend(sl)
	fmt.Println(rslt)

	if !IsSorted(rslt) {
		t.Errorf("got %v, but expect %v", IsSorted(rslt), true)
	}
}
