package linear_probing

import (
	"errors"
	"fmt"
	"strings"
)

type Fruit struct {
	Name  string
	Price int
}

type FruitCart [10]*Fruit // 🦊 for simplicity, use fixed size array to simplify hash function

func (fc *FruitCart) Put(name string, price int) error {
	if name == "" || price < 0 {
		panic("invalid key")
	}

	hash := hash(name)

	if fc[hash] != nil { // collision
		stopIndex := hash // stop probing at old position in case of full array
		if hash == len(fc)-1 {
			hash = 0 // try pushing at the first index
		} else {
			hash += 1 // try pushing after old position
		}
		for {
			if fc[hash] == nil { // found the empty index
				break // use this hash value
			}
			if hash == stopIndex { // array is full
				return errors.New("array is full, cannot put anymore")
			}
			hash = (hash + 1) % len(fc) // ... 8 -> 9 -> 0 -> 1 ...
		}
	}

	fc[hash] = &Fruit{
		name,
		price}
	return nil
}

func hash(key string) int {
	return len(key) % 10
}

func (fc FruitCart) Get(name string) (*Fruit, error) {
	hash, err := fc.IndexOf(name)
	if err != nil {
		return nil, errors.New("key not found")
	}
	return fc[hash], nil
}

func (fc FruitCart) IndexOf(name string) (int, error) {
	hash := hash(name)

	if fc[hash] != nil && fc[hash].Name == name {
		return hash, nil
	}

	stopIndex := hash
	if hash == len(fc)-1 {
		hash = 0
	} else {
		hash += 1
	}

	for {
		if fc[hash] == nil {
			return -1, errors.New("end of collision, item not found")
		}
		if fc[hash] != nil && fc[hash].Name == name {
			return hash, nil
		}
		if hash == stopIndex {
			return -1, errors.New("item not found")
		}
		hash = (hash + 1) % len(fc)
	}
}

// Remove
// rehashing when remove element that has collision
// once we hit a null, we know the value we're looking for isn't in the table. If we didn't do that,
// we could end up searching the entire array for values that aren't even in the table.
// That wouldn't matter for a small array, but for a hashtable that contains a large number of items,
// the performance of getting a value would degrade too much.
// The whole point of using a hashtable is that retrieving items is fast.
func (fc *FruitCart) Remove(name string) (*Fruit, error) {
	index, err := fc.IndexOf(name)
	if err != nil {
		return nil, errors.New("item not found")
	}
	rslt := fc[index]
	fc[index] = nil

	// rehashing entire array
	old := *fc // ⚠️ old := fc means store pointer of fc to old, clear fc will affect old
	for i := range fc {
		fc[i] = nil // clear array
	}
	for _, v := range old {
		if v != nil {
			err := fc.Put(v.Name, v.Price) // Put() will call hash() again
			if err != nil {
				panic(err)
			}
		}
	}

	return rslt, nil
}

func (fc FruitCart) LoadFactor() float32 {
	count := 0
	for _, v := range fc {
		if v != nil {
			count += 1
		}
	}
	return float32(count) / float32(len(fc))
}

func (fc FruitCart) Print() {
	sb := strings.Builder{}
	for _, v := range fc {
		if v == nil {
			sb.WriteString("nil ")
			continue
		}
		str := fmt.Sprintf("<%v %v> ", v.Name, v.Price)
		sb.WriteString(str)
	}
	fmt.Println(sb.String())
}
