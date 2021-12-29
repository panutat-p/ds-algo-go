package stack

import (
	"container/list"
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{list.New()}
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Size())
}

func TestStack_Push(t *testing.T) {
	s := Stack{list.New()}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Pop()
	s.Pop()
	s.Print()

	if s.Size() != 3 {
		t.Errorf("got %v, but expect %v", s.Size(), 3)
	}
}

func TestStack_Peek(t *testing.T) {
	s := Stack{list.New()}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()

	rst := s.Peek()

	if rst != 3 {
		t.Errorf("got %v, but expect %v", rst, 3)
	}
}
