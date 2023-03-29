package binary_search

import (
	"fmt"
	"testing"
)

func TestRecursiveSearchFor10(t *testing.T) {
	li := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("input:", li)
	rst := BinarySearchRecursive(li, 10)
	fmt.Println("output:", rst)

	if rst != true {
		t.Errorf("got %v, but expect %v", rst, 9)
	}
}

func TestRecursiveSearchFor12(t *testing.T) {
	li := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("input:", li)
	rst := BinarySearchRecursive(li, 12)
	fmt.Println("output:", rst)

	if rst != false {
		t.Errorf("got %v, but expect %v", rst, false)
	}
}

func TestRecursiveSearchFor1(t *testing.T) {
	li := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("input:", li)
	rst := BinarySearchRecursive(li, 1)
	fmt.Println("output:", rst)

	if rst != false {
		t.Errorf("got %v, but expect %v", rst, false)
	}
}
