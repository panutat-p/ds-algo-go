package chaining

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_LinkedList(t *testing.T) {
	li := list.New()
	li.PushBack(1)
	li.PushBack(2)
	li.PushBack(3)
	for e := li.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v -> ", e.Value)
	}
}

func TestFruitCart_Put(t *testing.T) {
	fc := New()
	fc.Put("apple", 15)
	fc.Put("banana", 8)
	fc.Put("carrot", 5)
	fc.Put("pear", 23)
	fc.Print()

	if fc.LoadFactor() != 0.4 {
		t.Errorf("got %v, but expect %v", fc.LoadFactor(), 0.4)
	}
}

func TestFruitCart_Remove_Collision_RemoveInLinkedList(t *testing.T) {
	fc := New()
	fc.Put("apple", 15)
	fc.Put("banana", 8)
	fc.Put("carrot", 5)
	fc.Put("pear", 23) // collide with apple
	fc.Print()

	fc.Remove("apple")
	fc.Print()

	if fc.LoadFactor() != 0.3 {
		t.Errorf("got %v, but expect %v", fc.LoadFactor(), 0.3)
	}

	if fc[7].Front().Value.(Fruit) != (Fruit{"pear", 23}) {
		t.Errorf("got %v, but expect %v", fc[7].Front().Value.(Fruit), Fruit{"pear", 23})
	}
}
