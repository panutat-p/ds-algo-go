package ordered_map

import (
	"fmt"
	"testing"
)

func TestLinkedMap(t *testing.T) {
	t.Run("7 elements", func(t *testing.T) {
		m := NewLinkedMap()
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)
		m.Put("d", 4)
		m.Put("e", 5)
		m.Put("f", 6)
		m.Put("g", 7)
		for _, e := range m.Keys() {
			fmt.Println(e, m.Get(e))
		}
		fmt.Println("m:", m)
		if m.Size() != 7 {
			t.Errorf("\nWant %d\nGot %d", 4, m.Size())
		}
	})
	t.Run("100 to 200", func(t *testing.T) {
		m := NewLinkedMap()
		for i := 100; i <= 200; i += 1 {
			m.Put(fmt.Sprintf("%d", i), i)
		}
		for _, e := range m.Keys() {
			fmt.Println(e, m.Get(e))
		}
		fmt.Println("m:", m)
		if m.Size() != 101 {
			t.Errorf("\nWant %d\nGot %d", 4, m.Size())
		}
	})
	t.Run("200 to 100", func(t *testing.T) {
		m := NewLinkedMap()
		for i := 200; i > 99; i -= 1 {
			m.Put(fmt.Sprintf("%d", i), i)
		}
		for _, e := range m.Keys() {
			fmt.Println(e, m.Get(e))
		}
		fmt.Println("m:", m)
		if m.Size() != 101 {
			t.Errorf("\nWant %d\nGot %d", 4, m.Size())
		}
	})
	t.Run("put then remove", func(t *testing.T) {
		m := NewLinkedMap()
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)
		m.Put("d", 4)
		m.Put("e", 5)
		m.Put("f", 6)
		m.Put("g", 7)
		fmt.Println("m before:", m)
		m.Remove("a")
		m.Remove("b")
		m.Remove("c")
		fmt.Println("m after:", m)
		if m.Size() != 4 {
			t.Errorf("\nWant %d\nGot %d", 4, m.Size())
		}
	})
	t.Run("put duplicated key", func(t *testing.T) {
		m := NewLinkedMap()
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)
		fmt.Println("m before:", m)
		m.Put("b", 4)
		fmt.Println("m after:", m)
		if m.Size() != 3 {
			t.Errorf("\nWant %d\nGot %d", 4, m.Size())
		}
	})
	t.Run("set keys", func(t *testing.T) {
		m := NewLinkedMap()
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)
		fmt.Println("m before:", m)
		m.Set("b", 20)
		m.Set("d", 4)
		fmt.Println("m after:", m)
		if m.Size() != 4 {
			t.Errorf("\nWant %d\nGot %d", 4, m.Size())
		}
	})
}
