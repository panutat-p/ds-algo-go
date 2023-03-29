package merge_sort

import (
	"fmt"
	"testing"

	"github.com/panutat-p/fiset-complete-ds-go/pkg"
)

func TestSplit(t *testing.T) {
	sl1, sl2 := SplitSlice([]int{0, 1, 2, 3, 4})
	fmt.Println(sl1, sl2)

	if len(sl1) != 2 {
		t.Errorf("got %v, but expect %v", len(sl1), 2)
	}

	if len(sl2) != 3 {
		t.Errorf("got %v, but expect %v", len(sl2), 3)
	}
}

func TestMerge(t *testing.T) {
	sl := MergeSlice([]int{1, 4, 6}, []int{2, 9})
	fmt.Println(sl)

	if len(sl) != 5 {
		t.Errorf("got %v, but expect %v", len(sl), 5)
	}

	if !pkg.IsSorted(sl) {
		t.Errorf("got %v, but expect %v", pkg.IsSorted(sl), true)
	}
}

func TestMergeSortSlice(t *testing.T) {
	sl1 := []int{34, 3, 4, 9, 0, 5, 65, 5, 41, 7}
	fmt.Println("sl1:", sl1)
	sl2 := MergeSortSlice(sl1)
	fmt.Println("sl2:", sl2)

	if sl2[0] > sl2[1] {
		t.Errorf("got %v, but expect %v", sl2[0] > sl2[1], false)
	}

	if !pkg.IsSorted(sl2) {
		t.Errorf("got %v, but expect %v", pkg.IsSorted(sl2), true)
	}
}
