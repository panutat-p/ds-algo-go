package bubble_sort

import (
	"reflect"
	"testing"
)

func TestPlainBubbleSort(t *testing.T) {
	tests := []struct {
		name  string
		given []int
		want  []int
	}{
		{
			name:  "1 to 5",
			given: []int{4, 1, 3, 5, 1, 2},
			want:  []int{1, 1, 2, 3, 4, 5},
		},
		{
			name:  "positive",
			given: []int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63},
			want:  []int{0, 3, 4, 5, 6, 7, 27, 30, 49, 51, 63, 100},
		},
		{
			name:  "negative",
			given: []int{14, 4, 26, -3, 0, 1, -34},
			want:  []int{-34, -3, 0, 1, 4, 14, 26},
		},
		{
			name:  "vary",
			given: []int{112, 12, 46, 3, 68, -3, 0, 1, 4, -49, 14, 26},
			want:  []int{-49, -3, 0, 1, 3, 4, 12, 14, 26, 46, 68, 112},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PlainBubbleSort(tt.given)
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("\nwant", tt.want, "\ngot ", got)
			}
		})
	}
}
