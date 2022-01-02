package priority_queue

import (
	"fmt"
	"testing"
)

func TestMaxHeap_HeapifyUp(t *testing.T) {
	h := MaxHeap{}
	fmt.Printf("%+v\n", h)
	h.Add(1)
	h.Add(2)
	h.Add(3)
	fmt.Printf("%+v\n", h)
	h.Add(100)
	fmt.Printf("%+v\n", h)

	if h.Size() != 4 {
		t.Errorf("got %v, but expect %v", h.Size(), 4)
	}
}

func TestMaxHeap_HeapifyDown(t *testing.T) {
	h := MaxHeap{}
	h.Add(100)
	h.Add(1)
	h.Add(2)
	h.Add(3)
	h.Add(4)
	fmt.Printf("%+v\n", h)

	extract, err := h.Poll()
	if err != nil {
		panic("extract failed")
	}
	fmt.Printf("%+v\n", h)

	if h.Size() != 4 {
		t.Errorf("got %v, but expect %v", h.Size(), 4)
	}

	if extract != 100 {
		t.Errorf("got %v, but expect %v", extract, 100)
	}
}
