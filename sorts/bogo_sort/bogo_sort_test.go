package bogo_sort

import (
	"fmt"
	"github.com/panutat-p/fiset-complete-ds-go/pkg"
	"testing"
)

func TestBogoSort(t *testing.T) {
	sl := []int{5, 7, 1, 21, 82, 44}
	sortedSl := BogoSort(sl)
	fmt.Println(sortedSl)
	if !pkg.IsSorted(sortedSl) {
		fmt.Println(sortedSl)
		t.Errorf("got %v, but expect %v", !pkg.IsSorted(sl), true)
	}
}
