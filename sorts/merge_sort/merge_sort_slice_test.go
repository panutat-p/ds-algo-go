package merge_sort

import (
	"slices"
	"testing"
)

func TestMergeSortSlice(t *testing.T) {
	tests := []struct {
		name  string
		given []int
	}{
		{
			name:  "1 to 5",
			given: []int{4, 1, 3, 5, 1, 2},
		},
		{
			name:  "positive",
			given: []int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63},
		},
		{
			name:  "negative",
			given: []int{14, 4, 26, -3, 0, 1, -34},
		},
		{
			name:  "pure negative",
			given: []int{-6, -5, -4, -11, -2},
		},
		{
			name:  "10 elements",
			given: []int{34, 3, 4, 9, 0, 5, 65, 5, 41, 7},
		},
		{
			name:  "12 elements",
			given: []int{112, 12, 46, 3, 68, -3, 0, 1, 4, -49, 14, 26},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSortSlice(tt.given)
			if !slices.IsSorted(got) {
				t.Error("\nWant", "sorted slice", "\nGot ", got)
			}
		})
	}
}
