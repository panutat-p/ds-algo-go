package chaining

import (
	"container/list"
	"errors"
	"fmt"
	"hash/fnv"
	"strings"
)

type Fruit struct {
	Name  string
	Price int
}

type FruitCart [10]*list.List // 🦊 for simplicity, use fixed size array to simplify hash function

func New() FruitCart {
	fc := FruitCart{}
	for i := range fc {
		fc[i] = list.New()
	}
	return fc
}

func (fc *FruitCart) Put(name string, price int) {
	if name == "" || price < 0 {
		panic("invalid key")
	}

	hash := hash(name)
	fmt.Println("name:", name, "hash:", hash)
	fc[hash].PushBack(Fruit{name, price}) // 🦊 linked list contains struct value
}

func hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32() % 10
}

func hash2(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % 10
}

func (fc *FruitCart) Remove(name string) (Fruit, error) {
	hash := hash(name)
	for e := fc[hash].Front(); e != nil; e = e.Next() {
		fruit := e.Value.(Fruit)
		if fruit.Name == name {
			fc[hash].Remove(e)
			return fruit, nil
		}
	}
	return Fruit{}, errors.New("item not found, cannot remove")
}

// Get
// This method cannot return *Fruit because linked list element stores Fruit in Value field
// read more at
// https://stackoverflow.com/questions/15518919/take-address-of-value-inside-an-interface
func (fc FruitCart) Get(name string) (Fruit, error) {
	hash := hash(name)
	for e := fc[hash].Front(); e != nil; e = e.Next() {
		if e.Value.(Fruit).Name == name {
			return e.Value.(Fruit), nil
		}
	}
	return Fruit{}, errors.New("item not found")
}

func (fc FruitCart) LoadFactor() float32 {
	count := 0
	for _, v := range fc {
		count += v.Len()
	}
	return float32(count) / float32(len(fc))
}

func (fc FruitCart) Print() {
	sb1 := strings.Builder{}
	sb2 := strings.Builder{}
	for _, v := range fc {
		if v.Len() != 0 {
			str1 := fmt.Sprintf("<[%v...] (%v)> ", v.Front().Value.(Fruit).Name, v.Len())
			sb1.WriteString(str1)

			for e := v.Front(); e != nil; e = e.Next() {
				str2 := fmt.Sprintf("<%v %v> ", e.Value.(Fruit).Name, e.Value.(Fruit).Price)
				sb2.WriteString(str2)
			}
		} else {
			size := fmt.Sprintf("<[]> ")
			sb1.WriteString(size)
		}
	}
	fmt.Println(sb1.String())
	fmt.Println(sb2.String())
}
