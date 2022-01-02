package priority_queue

import "errors"

type MaxHeap struct {
	sl []int
}

func (h *MaxHeap) Add(num int) {
	h.sl = append(h.sl, num)   // new num is on the last index
	h.heapifyUp(len(h.sl) - 1) // pass last index
}

// heapifyUp
// go from bottom to top
// index will be decreasing from lastIndex to zero
func (h *MaxHeap) heapifyUp(index int) {
	for h.sl[parent(index)] < h.sl[index] { // when index is 0, h.sl[0] < h.sl[0] must be false
		h.swap(parent(index), index)
		index = parent(index) // update current index to upper floor
	}
}

// Poll
// removes the root element
func (h *MaxHeap) Poll() (int, error) {
	if h.IsEmpty() {
		return -1, errors.New("heap is empty, cannot extract")
	}

	root := h.sl[0]
	h.sl[0] = h.sl[len(h.sl)-1] // assign new value to the root
	h.sl = h.sl[:len(h.sl)-1]
	h.heapifyDown(0) // pass the first index
	return root, nil
}

// heapifyDown ***
// go from top to bottom
// leftIndex and rightIndex will be increasing
func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := h.Size() - 1
	leftIndex := left(index)
	rightIndex := right(index)
	swappingIndex := 0
	for leftIndex < lastIndex {
		// there is a case when sub-tree is not a complete tree (single child)
		if leftIndex == lastIndex || h.sl[leftIndex] > h.sl[rightIndex] { // left child is larger
			swappingIndex = leftIndex
		} else { // right child is larger
			swappingIndex = rightIndex
		}

		if h.sl[index] < h.sl[swappingIndex] {
			h.swap(index, swappingIndex)
			index = swappingIndex
			leftIndex = left(index)
			rightIndex = right(index)
		} else { // found the right place, end heapify
			return
		}
	}
}

func (h MaxHeap) IsEmpty() bool {
	return h.Size() == 0
}

func (h MaxHeap) Size() int {
	return len(h.sl)
}

// parent
// pass left child index >> divide evenly
// pass right child index >> round the result
func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (2 * i) + 1
}

func right(i int) int {
	return (2 * i) + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.sl[i1], h.sl[i2] = h.sl[i2], h.sl[i1]
}
