package bubble_sort

import (
	"slices"
	"testing"
)

func TestOptimizedSort(t *testing.T) {
	t.Run("1 to 5", func(t *testing.T) {
		sl := []int{1, 2, 3, 4, 5}
		OptimizedSort(sl)
		if !slices.IsSorted(sl) {
			t.Error("\nslice is not sorted", "\nsl ", sl)
		}
	})
	t.Run("positive", func(t *testing.T) {
		sl := []int{0, 3, 4, 5, 6, 7, 27, 30, 49, 51, 63, 100}
		OptimizedSort(sl)
		if !slices.IsSorted(sl) {
			t.Error("\nslice is not sorted", "\nsl ", sl)
		}
	})
	t.Run("negative", func(t *testing.T) {
		sl := []int{-34, -3, 0, 1, 4, 14, 26}
		OptimizedSort(sl)
		if !slices.IsSorted(sl) {
			t.Error("\nslice is not sorted", "\nsl ", sl)
		}
	})
	t.Run("vary", func(t *testing.T) {
		sl := []int{-49, -3, 0, 1, 3, 4, 12, 14, 26, 46, 68, 112}
		OptimizedSort(sl)
		if !slices.IsSorted(sl) {
			t.Error("\nslice is not sorted", "\nsl ", sl)
		}
	})
}
