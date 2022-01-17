package linear_probing

import (
	"testing"
)

func TestFruitCart_Push(t *testing.T) {
	fc := FruitCart{}
	fc.Put("banana", 10)
	fc.Put("apple", 15)
	fc.Put("watermelon", 18)
	fc.Print()

	if fc.LoadFactor() != 0.3 {
		t.Errorf("got %v, but expect %v", fc.LoadFactor(), 0.3)
	}
}

func TestFruitCart_Push_collision(t *testing.T) {
	fc := FruitCart{}
	fc.Put("banana", 10)
	fc.Put("BANANA", 12)
	fc.Put("apple", 15)
	fc.Put("watermelon", 18)
	fc.Put("WATERMELON", 20)
	fc.Print()

	rslt, err := fc.Get("BANANA")
	if err != nil {
		t.Errorf("got %v, but expect %v", err, nil)
	}

	if rslt != (Fruit{"BANANA", 12}) {
		t.Errorf("got %v, but expect %v", rslt, Fruit{"BANANA", 12})
	}
}

func TestFruitCart_Find_collision(t *testing.T) {
	fc := FruitCart{}
	fc.Put("banana", 10)
	fc.Put("BANANA", 12)
	fc.Put("apple", 15)
	fc.Put("watermelon", 18)
	fc.Put("WATERMELON", 20)
	fc.Print()

	rslt, err := fc.Get("BANANA")

	if err != nil {
		t.Errorf("got %v, but expect %v", err, nil)
	}

	if rslt != (Fruit{"BANANA", 12}) {
		t.Errorf("got %v, but expect %v", rslt, Fruit{"BANANA", 12})
	}
}

func TestFruitCart_Remove_rehashing(t *testing.T) {
	fc := FruitCart{}
	fc.Put("banana", 10)
	fc.Put("BANANA", 12)
	fc.Put("apple", 15)
	fc.Put("watermelon", 18)
	fc.Put("WATERMELON", 20)

	fc.Print()
	rslt, err := fc.Remove("watermelon") // WATERMELON will be relocated to index 0
	fc.Print()

	if err != nil {
		t.Errorf("got %v, but expect %v", err, nil)
	}

	if rslt != (Fruit{"watermelon", 18}) {
		t.Errorf("got %v, but expect %v", rslt, Fruit{"watermelon", 18})
	}

	if fc[0] != (Fruit{"WATERMELON", 20}) {
		t.Errorf("got %v, but expect %v", fc[0], Fruit{"WATERMELON", 20})
	}
}
