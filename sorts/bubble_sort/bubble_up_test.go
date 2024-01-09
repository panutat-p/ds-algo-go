package bubble_sort

import (
	"reflect"
	"testing"
)

func TestBubbleUp(t *testing.T) {
	t.Run("1 to 5", func(t *testing.T) {
		want := []int{1, 1, 2, 3, 4, 5}
		got := BubbleUp([]int{4, 1, 3, 5, 1, 2})
		if !reflect.DeepEqual(want, got) {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("positive", func(t *testing.T) {
		want := []int{0, 3, 4, 5, 6, 7, 27, 30, 49, 51, 63, 100}
		got := BubbleUp([]int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63})
		if !reflect.DeepEqual(want, got) {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("negative", func(t *testing.T) {
		want := []int{-34, -3, 0, 1, 4, 14, 26}
		got := BubbleUp([]int{14, 4, 26, -3, 0, 1, -34})
		if !reflect.DeepEqual(want, got) {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("vary", func(t *testing.T) {
		want := []int{-49, -3, 0, 1, 3, 4, 12, 14, 26, 46, 68, 112}
		got := BubbleUp([]int{112, 12, 46, 3, 68, -3, 0, 1, 4, -49, 14, 26})
		if !reflect.DeepEqual(want, got) {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
}
