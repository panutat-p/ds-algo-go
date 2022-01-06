package insertion_sort

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

func TestInsertionSort(t *testing.T) {
	sl := []int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63}
	InsertionSort(sl)
	fmt.Println(sl)

	if !IsSorted(sl) {
		t.Errorf("got %v, but expect %v", IsSorted(sl), true)
	}
}

func TestInsertionSort2(t *testing.T) {
	sl := []int{1, 5, 6, 27, 100, 100, 250}
	InsertionSort(sl)
	fmt.Println(sl)

	if !IsSorted(sl) {
		t.Errorf("got %v, but expect %v", IsSorted(sl), true)
	}
}

func TestRecursiveInsertionSort(t *testing.T) {
	sl := []int{30, 5, 6, 27, 100, 3, 51, 7, 49, 0, 4, 63}
	RecursiveInsertionSort(sl, len(sl))
	fmt.Println(sl)

	if !IsSorted(sl) {
		t.Errorf("got %v, but expect %v", IsSorted(sl), true)
	}
}
