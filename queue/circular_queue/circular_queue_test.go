package circular_queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue_Size(t *testing.T) {
	cq := New()
	for i := 0; i < 10; i += 1 {
		fmt.Printf("%+v\n", cq)
		fmt.Println("Size:", cq.Size())
		cq.Enqueue(i)
	}
	fmt.Println(cq.Dequeue())
	fmt.Println(cq.Dequeue())
	fmt.Printf("%+v\n", cq)
	fmt.Println("Size:", cq.Size())

	cq.Print()

	if cq.Size() != 8 {
		t.Errorf("got %v, but expect %v", cq.Size(), 8)
	}
}

func TestCircularQueue_reSizeSlice(t *testing.T) {
	cq := New()
	for i := 0; i < 6; i += 1 {
		cq.Enqueue(i)
	}
	fmt.Printf("%+v\n", cq)
	fmt.Println("Size:", cq.Size())
	cq.Dequeue()
	cq.Dequeue()
	cq.Dequeue()
	cq.Dequeue()
	fmt.Printf("%+v\n", cq)
	fmt.Println("Size:", cq.Size())
	for i := 6; i < 13; i += 1 {
		cq.Enqueue(i)
		fmt.Printf("%+v\n", cq)
		fmt.Println("Size:", cq.Size())
	}

	cq.Print()

	if cq.Size() != 9 {
		t.Errorf("got %v, but expect %v", cq.Size(), 9)
	}
}

func TestCircularQueue_Dequeue(t *testing.T) {
	cq := New()
	for i := 0; i < 6; i += 1 {
		cq.Enqueue(i)
	}
	cq.Dequeue()
	cq.Dequeue()
	cq.Dequeue()
	cq.Dequeue()
	for i := 6; i < 10; i += 1 {
		cq.Enqueue(i)
	}
	fmt.Printf("%+v\n", cq)
	fmt.Println("Size:", cq.Size())
	cq.Dequeue()
	cq.Dequeue()
	cq.Dequeue()
	fmt.Printf("%+v\n", cq)
	fmt.Println("Size:", cq.Size())
	cq.Print()

	if cq.Size() != 3 {
		t.Errorf("got %v, but expect %v", cq.Size(), 3)
	}

	cq.Dequeue()
	cq.Dequeue()
	cq.Dequeue()
	cq.Enqueue(100)
	fmt.Printf("%+v\n", cq)
	fmt.Println("Size:", cq.Size())
	cq.Print()

	if cq.Size() != 1 {
		t.Errorf("got %v, but expect %v", cq.Size(), 1)
	}
}
