package dynamic_array

import (
	"fmt"
	"testing"
)

func TestMakeDynamicArray(t *testing.T) {
	da := MakeDynamicArray()
	fmt.Println(da)
	fmt.Println(da.Arr)
	da.Append(1)
	fmt.Println(da)
	fmt.Println(da.Arr)
	da.Append(2)
	fmt.Println(da)
	fmt.Println(da.Arr)
	if da.Size != 2 {
		t.Errorf("got %v, but expect %v", da.Size, 2)
	}

	if da.Capacity != 16 {
		t.Errorf("got %v, but expect %v", da.Capacity, 16)
	}
}

func TestDynamicArray_Pop(t *testing.T) {
	da := MakeDynamicArray()
	da.Append(1)
	da.Append(2)
	da.Append(3)
	da.Append(4)
	fmt.Println(da)
	rst, err := da.Pop(2)
	fmt.Println(da)
	//fmt.Println(da.Arr[3])

	if err != nil {
		t.Errorf("got %v, but expect %v", err, nil)
	}

	if rst != 3 {
		t.Errorf("got %v, but expect %v", rst, 3)
	}
}

func TestDynamicArray_Remove(t *testing.T) {
	da := MakeDynamicArray()
	da.Append(1)
	da.Append(2)
	da.Append(3)
	da.Append(4)
	fmt.Println(da)
	rst := da.Remove(2)
	fmt.Println(da)

	if !rst {
		t.Errorf("got %v, but expect %v", rst, true)
	}

	if da.Size != 3 {
		t.Errorf("got %v, but expect %v", da.Size, 3)
	}
}
