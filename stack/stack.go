package stack

import (
	"container/list"
	"fmt"
	"strings"
)

type Stack struct {
	li *list.List
}

func New() *Stack {
	return &Stack{list.New()}
}

func (s Stack) Print() {
	var forward strings.Builder
	current := s.li.Front()
	for {
		if current == nil {
			forward.WriteString("nil")
			break
		}
		tmp := fmt.Sprintf("%v -> ", current.Value)
		forward.WriteString(tmp)
		current = current.Next()
	}
	fmt.Println(forward.String())
}

func (s Stack) Size() int {
	return s.li.Len()
}

func (s Stack) IsEmpty() bool {
	return s.li.Len() == 0
}

func (s Stack) Peek() interface{} {
	return s.li.Back().Value
}

func (s *Stack) Push(v interface{}) {
	s.li.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	ele := s.li.Back()
	s.li.Remove(ele)
	return ele.Value
}
