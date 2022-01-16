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
	if rslt.Name != "BANANA" {
		t.Errorf("got %v, but expect %v", rslt, "BANANA")
	}

	if rslt.Price != 12 {
		t.Errorf("got %v, but expect %v", rslt.Price, 12)
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

	if rslt.Name != "BANANA" {
		t.Errorf("got %v, but expect %v", rslt, "BANANA")
	}

	if rslt.Price != 12 {
		t.Errorf("got %v, but expect %v", rslt.Price, 12)
	}
}

func TestFruitCart_Remove(t *testing.T) {
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

	if rslt.Name != "watermelon" {
		t.Errorf("got %v, but expect %v", rslt, "watermelon")
	}

	if rslt.Price != 18 {
		t.Errorf("got %v, but expect %v", rslt.Price, 18)
	}
}
