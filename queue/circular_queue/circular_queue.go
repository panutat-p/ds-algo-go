package circular_queue

import (
	"errors"
	"fmt"
)

/*
https://www.programiz.com/dsa/circular-queue
*/

type CircularQueue struct {
	sl         []interface{}
	frontIndex int // first element
	rearIndex  int // behind the last element
	isWrapped  bool
}

func New() *CircularQueue {
	sl := make([]interface{}, 4)
	return &CircularQueue{sl, 0, 0, false}
}

func (cq CircularQueue) Peek() interface{} {
	return cq.sl[cq.frontIndex]
}

func (cq *CircularQueue) Enqueue(v interface{}) {
	// in case of no Dequeue() call
	// we cannot wrap the slice to replace the value at first index
	// so, we resize the slice
	if cq.Size() == len(cq.sl)-1 { // 🦄 we have to resize the array when there's only one empty position at the last index
		cq.resizeSlice()
	}

	cq.sl[cq.rearIndex] = v
	if cq.rearIndex == len(cq.sl)-1 { // reach at the last element
		cq.rearIndex = 0 // 🦄 wrap
		cq.isWrapped = true
	} else {
		cq.rearIndex += 1 // right shift to empty element
	}
}

func (cq *CircularQueue) Dequeue() (interface{}, error) {
	if cq.Size() == 0 {
		return nil, errors.New("cannot dequeue, queue is empty")
	}

	rslt := cq.sl[cq.frontIndex]
	cq.sl[cq.frontIndex] = nil
	cq.frontIndex += 1

	// 🦊 handle boundary cases
	if cq.Size() == 0 { // reset queue 👍🏻
		// this happens when frontIndex hit the rearIndex
		cq.frontIndex = 0
		cq.rearIndex = 0
		cq.isWrapped = false
	} else if cq.isWrapped && cq.frontIndex == len(cq.sl) {
		// this can happen when slice was wrapped, if not, frontIndex will hit rearIndex before hit the length
		cq.frontIndex = 0
	}

	return rslt, nil
}

func (cq CircularQueue) Size() int {
	if cq.frontIndex <= cq.rearIndex {
		return cq.rearIndex - cq.frontIndex
	} else { // 🦄 already wrapped
		return len(cq.sl) + (cq.rearIndex - cq.frontIndex) // total length - number of empty elements
	}
}

// resize and unwrap the slice if slice was wrapped
// [4,5,1,2,3] >> [1,2,3,4,5,...]
//    ^ ^          ^       ^
func (cq *CircularQueue) resizeSlice() {
	sl := make([]interface{}, 2*len(cq.sl))
	if cq.isWrapped {
		index := 0
		for i := cq.frontIndex; i < len(cq.sl); i += 1 {
			sl[index] = cq.sl[i]
			index += 1
		}
		for i := 0; i < cq.rearIndex; i += 1 {
			sl[index] = cq.sl[i]
			index += 1
		}
		cq.sl = sl
		cq.isWrapped = false
		cq.frontIndex = 0
		cq.rearIndex = index
	} else {
		for i, v := range cq.sl {
			sl[i] = v
		}
		cq.sl = sl
	}
}

func (cq CircularQueue) Print() {
	if cq.isWrapped {
		for i := cq.frontIndex; i < len(cq.sl); i += 1 {
			fmt.Printf("%v -> ", cq.sl[i])
		}
		for i := 0; i < cq.rearIndex; i += 1 {
			fmt.Printf("%v -> ", cq.sl[i])
		}
	} else {
		for i := 0; i < cq.rearIndex; i += 1 {
			fmt.Printf("%v -> ", cq.sl[i])
		}
	}
	fmt.Println()
}
