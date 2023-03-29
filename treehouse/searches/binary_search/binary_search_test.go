package binary_search

import (
	"fmt"
	"testing"
)

func TestSearchFor10(t *testing.T) {
	li := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("input", li)
	index := BinarySearch(li, 10)
	fmt.Println("output", index)

	if index != 9 {
		t.Errorf("got %v, but expect %v", index, 9)
	}
}

func TestSearchFor12(t *testing.T) {
	li := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("input", li)
	index := BinarySearch(li, 12)
	fmt.Println("output", index)

	if index != -1 {
		t.Errorf("got %v, but expect %v", index, -1)
	}
}

func TestSearchFor1(t *testing.T) {
	li := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("input", li)
	index := BinarySearch(li, 1)
	fmt.Println("output", index)

	if index != -1 {
		t.Errorf("got %v, but expect %v", index, -1)
	}
}
