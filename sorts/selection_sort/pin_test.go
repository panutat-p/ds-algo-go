package selection_sort

import (
	"testing"

	"github.com/panutat-p/fiset-complete-ds-go/pkg"
)

func TestSort1_already_sorted(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := Sort1(sl)

	if !pkg.IsSorted(sl) {
		t.Error("The slice is not sorted:", got)
	}
}

func TestSort1_positive_numbers(t *testing.T) {
	sl := []int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63}
	got := Sort1(sl)

	if !pkg.IsSorted(sl) {
		t.Error("The slice is not sorted:", got)
	}
}
