package selection_sort

import (
	"fmt"
	"testing"

	"github.com/panutat-p/fiset-complete-ds-go/pkg"
)

func TestSort3_positive_numbers(t *testing.T) {
	sl := []int{5, 3, 7, 0, 14, 8, 1, 13, 21, 82, 44}
	sortedSl := Sort3(sl)
	fmt.Println(sortedSl)

	if !pkg.IsSorted(sortedSl) {
		fmt.Println(sortedSl)
		t.Errorf("got %v, but expect %v", pkg.IsSorted(sortedSl), true)
	}
}
