package bucket_sort

import (
	"fmt"
	"testing"
)

func IsSorted(sl [10]int) bool {
	for i := 0; i < len(sl)-1; i += 1 {
		if sl[i] > sl[i+1] {
			return false
		}
	}
	return true
}

func TestBucketSort(t *testing.T) {
	arr := [10]int{3, 5, 6, 3, 7, 4, 3, 1, 2, 0}
	BucketSort(&arr, 1)
	fmt.Println(arr)

	if !IsSorted(arr) {
		t.Errorf("got %v, but expect %v", IsSorted(arr), true)
	}
}

func TestBucketSort2(t *testing.T) {
	arr := [10]int{31, 52, 65, 37, 74, 54, 73, 81, 12, 20}
	BucketSort(&arr, 2)
	fmt.Println(arr)

	if !IsSorted(arr) {
		t.Errorf("got %v, but expect %v", IsSorted(arr), true)
	}
}
