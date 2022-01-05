package counting_sort

import (
	"fmt"
	"testing"
)

func TestStableCountingSort(t *testing.T) {
	sl := []int{5, 2, 5, 4, 1, 3}

	rslt := StableCountingSort(sl, 1, 5)
	fmt.Println(rslt)

	if !IsSorted(rslt) {
		t.Errorf("got %v, but expect %v", IsSorted(rslt), true)
	}
}

func TestStableCountingSort2(t *testing.T) {
	sl := []int{4, 4, 4, 3, 3, 2, 1, 0}

	rslt := StableCountingSort(sl, 0, 4)
	fmt.Println(rslt)

	if !IsSorted(rslt) {
		t.Errorf("got %v, but expect %v", IsSorted(rslt), true)
	}
}
