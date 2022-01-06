package radix_sort

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

func TestGetDigit(t *testing.T) {
	rslt := getDigit(1024, 1, 10)
	fmt.Println("rslt:", rslt)
}

func TestRadixSort(t *testing.T) {
	sl := []int{5, 2, 5, 4, 1, 3}

	rslt := RadixSort(sl, 10, 1)
	fmt.Println(rslt)

	if !IsSorted(rslt) {
		t.Errorf("got %v, but expect %v", IsSorted(rslt), true)
	}
}

func TestRadixSort2(t *testing.T) {
	sl := []int{51, 22, 53, 48, 19, 30}

	rslt := RadixSort(sl, 10, 2)
	fmt.Println(rslt)

	if !IsSorted(rslt) {
		t.Errorf("got %v, but expect %v", IsSorted(rslt), true)
	}
}
