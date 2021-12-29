package stack

import (
	"container/list"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{list.New()}
	_, err := q.Peek()

	if err == nil {
		t.Errorf("got %v, but expect %v", err, "queue is empty")
	}
}

func TestQueue_Enqueue(t *testing.T) {
	q := Queue{list.New()}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Print()
	q.Dequeue()
	q.Print()

	if q.Size() != 2 {
		t.Errorf("got %v, but expect %v", q.Size(), 3)
	}
}
