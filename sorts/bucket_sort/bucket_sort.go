package bucket_sort

import (
	"container/list"
	"math"
)

// BucketSort
// not inplace
// use bubble sort
// size of input array is fixed to 10
func BucketSort(arr *[10]int, width int) {
	var buckets [10]*list.List
	for i := range buckets {
		buckets[i] = list.New()
	}

	for _, v := range arr {
		index := hash(v, width)
		buckets[index].PushBack(v)
	}

	for _, v := range buckets {
		sortLinkedList(v)
	}

	index := 0
	for _, v := range buckets {
		for e := v.Front(); e != nil; e = e.Next() {
			arr[index] = e.Value.(int)
			index += 1
		}
	}
}

// applicable for array with size 10, fixed size of width of each element
// 4/1 = 4, 4%10 = 4
// 14/10 = 4, 4%10 = 4
func hash(ele int, width int) uint32 {
	return uint32(ele/int(math.Pow(10, float64(width-1)))) % 10
}

// https://go.dev/play/p/Cyp5XnZqgjN
func sortLinkedList(li *list.List) *list.List {
	current := li.Front()
	if li.Front() == nil {
		return nil
	} else {
		for current != nil {
			index := current.Next()
			for index != nil {
				if current.Value.(int) > index.Value.(int) {
					current.Value, index.Value = index.Value, current.Value
				}
				index = index.Next()
			}
			current = current.Next()
		}
	}
	return li
}
