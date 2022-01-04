package counting_sort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	sl := []int{5, 2, 5, 4, 1, 3}

	CountingSort(sl, 1, 5)
	fmt.Println(sl)

	if !IsSorted(sl) {
		t.Errorf("got %v, but expect %v", IsSorted(sl), true)
	}
}

func IsSorted(sl []int) bool {
	for i := 0; i < len(sl)-1; i += 1 {
		if sl[i] > sl[i+1] {
			return false
		}
	}
	return true
}
