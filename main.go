package main

import "fmt"

func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

type Stack []int

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}

func (s *Stack) Pop() int {
	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e
}
