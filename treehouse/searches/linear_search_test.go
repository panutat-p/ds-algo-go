package searches

import (
	"fmt"
	"testing"
)

func TestSearchFor6(t *testing.T) {
	li := []int{1, 3, 4, 6, 7}
	fmt.Println("input:", li)
	index := LinearSearch(li, 6)
	fmt.Println("output:", index)

	if index != 3 {
		t.Errorf("got %v, but expect %v", index, 3)
	}
}

func TestSearchNotFound(t *testing.T) {
	li := []int{1, 3, 4, 6, 7, 4, 3, 5, 6, 9}
	fmt.Println("input:", li)
	index := LinearSearch(li, 12)
	fmt.Println("output:", index)

	if index != -1 {
		t.Errorf("got %v, but expect %v", index, -1)
	}
}
