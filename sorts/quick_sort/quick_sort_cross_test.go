package quick_sort

import (
	"fmt"
	"slices"
	"testing"
)

func TestQuickSortCross(t *testing.T) {
	t.Run("12 elements", func(t *testing.T) {
		sl := []int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63}
		QuickSortCross(sl, 0, len(sl))
		fmt.Println(sl)
		if !slices.IsSorted(sl) {
			t.Error("\nslice is not sorted\ngot", sl)
		}
	})

	t.Run("sorted input", func(t *testing.T) {
		sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		QuickSortCross(sl, 0, len(sl))
		fmt.Println(sl)
		if !slices.IsSorted(sl) {
			t.Error("\nslice is not sorted\ngot", sl)
		}
	})
}
