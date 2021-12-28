package dynamic_array

import (
	"errors"
)

type dynamicArray struct {
	Arr      []int
	Size     int // length that user thinks array is
	Capacity int // actual size
}

// MakeDynamicArray
// constructor
func MakeDynamicArray() dynamicArray {
	initialCapacity := 16
	sl := make([]int, 0, initialCapacity)
	return dynamicArray{sl, 0, initialCapacity}
}

func (d dynamicArray) IsEmpty() bool {
	return d.Size == 0
}

func (d dynamicArray) Get(index int) int {
	return d.Arr[index]
}

func (d *dynamicArray) Set(index int, ele int) {
	(*d).Arr[index] = ele
}

func (d *dynamicArray) Clear() {
	for i := 0; i < (*d).Size; i += 1 {
		(*d).Arr[i] = 0
	}
}

// Append
// resize array if Size reach Capacity
// allow using d instead of *d in the function body without the explicit dereference
func (d *dynamicArray) Append(ele int) {
	if d.Size+1 > d.Capacity {
		d.Capacity *= 2                        // grow exponential
		arr := make([]int, d.Size, d.Capacity) // make new array with double size
		for i := 0; i < d.Size; i += 1 {
			arr[i] = d.Arr[i]
		}
		d.Arr = arr
	}

	d.Arr = d.Arr[0 : d.Size+1]
	d.Arr[d.Size] = ele // old Size is the last index
	d.Size += 1
}

func (d *dynamicArray) Pop(index int) (int, error) {
	if index > d.Size || index < 0 {
		return -1, errors.New("index out of range")
	}

	rst := d.Arr[index]

	//d.Arr = append(d.Arr[:index], d.Arr[index+1:]...)

	arr := make([]int, d.Size-1, d.Capacity-1)
	i, j := 0, 0
	for i < d.Size {
		if i == index {
			i += 1
		} else {
			arr[j] = d.Arr[i]
			i += 1
			j += 1
		}
	}
	d.Arr = arr
	d.Size -= 1
	d.Capacity -= 1
	return rst, nil
}

func (d *dynamicArray) Remove(num int) bool {
	var index int
	for i, v := range d.Arr {
		if v == num {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	_, err := d.Pop(index)
	if err != nil {
		return false
	}
	return true
}
