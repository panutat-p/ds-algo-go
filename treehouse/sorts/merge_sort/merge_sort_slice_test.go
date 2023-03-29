package merge_sort

import (
	"fmt"
	"testing"

	"github.com/panutat-p/fiset-complete-ds-go/treehouse/sorts/bogo_sort"
)

func TestSplit(t *testing.T) {
	sl1, sl2 := Split([]int{0, 1, 2, 3, 4})
	fmt.Println(sl1, sl2)

	if len(sl1) != 2 {
		t.Errorf("got %v, but expect %v", len(sl1), 2)
	}

	if len(sl2) != 3 {
		t.Errorf("got %v, but expect %v", len(sl2), 3)
	}
}

func TestMerge(t *testing.T) {
	sl := Merge([]int{1, 4, 6}, []int{2, 9})
	fmt.Println(sl)

	if len(sl) != 5 {
		t.Errorf("got %v, but expect %v", len(sl), 5)
	}

	if !bogo_sort.IsSorted(sl) {
		t.Errorf("got %v, but expect %v", bogo_sort.IsSorted(sl), true)
	}
}

func TestMergeSort(t *testing.T) {
	sl1 := []int{34, 3, 4, 9, 0, 5, 65, 5, 41, 7}
	fmt.Println("sl1:", sl1)
	sl2 := MergeSort(sl1)
	fmt.Println("sl2:", sl2)

	if sl2[0] > sl2[1] {
		t.Errorf("got %v, but expect %v", sl2[0] > sl2[1], false)
	}

	if !bogo_sort.IsSorted(sl2) {
		t.Errorf("got %v, but expect %v", bogo_sort.IsSorted(sl2), true)
	}
}
