package merge_sort

import (
	"slices"
	"testing"
)

func TestMergeSortSlice(t *testing.T) {
	t.Run("10 elements", func(t *testing.T) {
		given := []int{34, 3, 4, 9, 0, 5, 65, 5, 41, 7}
		got := MergeSortSlice(given)
		if !slices.IsSorted(got) {
			t.Errorf("\nGot %v\nWant sorted slice", got)
		}
	})
	t.Run("positives", func(t *testing.T) {
		given := []int{7, 3, 5, 4, 9, 1}
		got := MergeSortSlice(given)
		if !slices.IsSorted(got) {
			t.Errorf("\nGot %v\nWant sorted slice", got)
		}
	})
	t.Run("negatives", func(t *testing.T) {
		given := []int{-6, -5, -4, -11, -2}
		got := MergeSortSlice(given)
		if !slices.IsSorted(got) {
			t.Errorf("\nGot %v\nWant sorted slice", got)
		}
	})
}
