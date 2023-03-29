package bogo_sort

import (
	"fmt"
	"testing"
)

func TestBogoSort(t *testing.T) {
	sl := []int{5, 3, 7, 0, 14, 8, 1, 13, 21, 82, 44}
	sortedSl := BogoSort(sl)
	fmt.Println(sortedSl)
	if !IsSorted(sortedSl) {
		fmt.Println(sortedSl)
		t.Errorf("got %v, but expect %v", !IsSorted(sl), true)
	}
}
