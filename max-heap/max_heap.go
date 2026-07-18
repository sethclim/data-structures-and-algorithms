package maxHeap

import (
	"fmt"
)

type MaxHeap struct {
	arr []int
}

func (h *MaxHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MaxHeap) BuildHeap(arr []int) {
	h.arr = append([]int(nil), arr...)

	lastNonLeaf := len(h.arr)/2 - 1

	for i := lastNonLeaf; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MaxHeap) ExtractMax() (int, error) {

	if len(h.arr) == 0 {
		return -1, fmt.Errorf("cannot extract from empty heap")
	}

	oldMax := h.arr[0]

	last := len(h.arr) - 1
	h.arr[0] = h.arr[last]
	h.arr = h.arr[:last]

	if len(h.arr) > 0 {
		h.heapifyDown(0)
	}

	return oldMax, nil
}

func (h *MaxHeap) Peek() (int, error) {
	if len(h.arr) == 0 {
		return -1, fmt.Errorf("cannot peek an empty heap")
	}

	return h.arr[0], nil
}

func (h *MaxHeap) heapifyDown(i int) {
	left := 2*i + 1
	right := 2*i + 2

	var largest int

	if left < len(h.arr) && h.arr[left] > h.arr[i] {
		largest = left
	} else {
		largest = i
	}

	if right < len(h.arr) && h.arr[right] > h.arr[largest] {
		largest = right
	}

	if largest != i {
		h.swap(i, largest)
		h.heapifyDown(largest)
	}
}

func (h *MaxHeap) heapifyUp(i int) {
	for i > 0 {
		p := (i - 1) / 2

		if h.arr[i] <= h.arr[p] {
			break
		}

		h.swap(i, p)

		i = p
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
