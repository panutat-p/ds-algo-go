package singly_linked_list_generic

import (
	"fmt"
	"testing"
)

func TestGenericLinkedList(t *testing.T) {
	t.Run("4 elements", func(t *testing.T) {
		m := NewLinkedList[int]()
		m.Add(1)
		m.Add(2)
		m.Add(3)
		m.Add(4)
		fmt.Println("m:", m)
		for e := m.Front(); e != nil; e = e.Next {
			fmt.Println(e.Value)
		}
		if m.Size() != 4 {
			t.Errorf("\nWant %d\nGot %d", 4, m.Size())
		}
	})
}
