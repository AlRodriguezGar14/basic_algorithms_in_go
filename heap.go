package main

import (
	"fmt"
)

// In the heap the values are ordered such a way that they
// match with the indexes of an array.
//
//	     (1)A
//	(2)B       (3)C
//
// (4)D (5)E   (6)F (7)G
//
// [A, B, C, D, E, F, G]

type minHeap struct {
	Length int
	Data   []int
}

func heapConstructor() *minHeap {
	return &minHeap{Length: 0, Data: []int{}}
}

func (h *minHeap) InsertHeap(value int) {
	if h.Length < cap(h.Data) {
		h.Data[h.Length] = value
	} else {
		newData := make([]int, h.Length+1)
		copy(newData, h.Data)
		newData[h.Length] = value
		h.Data = newData
	}
	h.heapifyUp(h.Length)
	h.Length++
}

// Delete the one on top, move the rest, re-balance the tree
func (h *minHeap) DeleteHeap() (int, error) {
	out := h.Data[0]
	h.Length--
	if h.Length == 0 {
		h.Data = []int{}
		return out, nil
	}
	h.Data[0] = h.Data[h.Length]
	h.heapifyDown(0)
	return out, nil
}

func (h *minHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parentIdx := h.parent(idx)
	parentValue := h.Data[parentIdx]
	v := h.Data[idx]

	if parentValue > v {
		h.Data[idx], h.Data[parentIdx] = parentValue, v
		h.heapifyUp(parentIdx)
	}

}

func (h *minHeap) heapifyDown(idx int) {
	if idx >= h.Length {
		return
	}

	leftIdx := h.leftChild(idx)

	if leftIdx >= h.Length {
		return
	}

	rightIdx := h.rightChild(idx)
	leftValue := h.Data[leftIdx]
	rightValue := h.Data[rightIdx]
	v := h.Data[idx]

	// by design, the right is the small option
	if leftValue > rightValue && v > rightValue {
		h.Data[idx] = rightValue
		h.Data[rightIdx] = v
		h.heapifyDown(rightIdx)
	} else if leftValue < rightValue && v < rightValue {
		h.Data[idx] = leftValue
		h.Data[leftIdx] = v
		h.heapifyDown(leftIdx)
	}
	return
}

func (h *minHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *minHeap) leftChild(idx int) int {
	return (idx * 2) + 1
}

func (h *minHeap) rightChild(idx int) int {
	return (idx * 2) + 2
}

func testHeap() {
	heap := heapConstructor()
	heap.InsertHeap(10)
	// heap.InsertHeap(20)
	// heap.InsertHeap(24)
	// heap.InsertHeap(28)
	// heap.InsertHeap(27)
	// heap.InsertHeap(2)
	fmt.Println(heap)
	heap.DeleteHeap()
	fmt.Println(heap)
}
