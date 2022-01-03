package merge_sort

import (
	"fmt"
	"testing"
)

func IsSorted(sl []int) bool {
	for i := 0; i < len(sl)-1; i += 1 {
		if sl[i] > sl[i+1] {
			return false
		}
	}
	return true
}

func TestMergeSort(t *testing.T) {
	sl := []int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63}
	rslt := MergeSort(sl)
	fmt.Println(rslt)

	if !IsSorted(rslt) {
		t.Errorf("got %v, but expect %v", IsSorted(rslt), true)
	}
}

func TestMergeSort2(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rslt := MergeSort(sl)
	fmt.Println(rslt)

	if !IsSorted(rslt) {
		t.Errorf("got %v, but expect %v", IsSorted(rslt), true)
	}
}
