package main

import "fmt"

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

}

func (h *minHeap) DeleteHeap() int {
	return 0
}

func (h *minHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parent := h.parent(idx)
	parentValue := h.Data[parent]
	v := h.Data[idx]

	if parentValue > v {
		h.Data[idx], h.Data[parent] = parentValue, v
		h.heapifyUp(parent)
	}

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
	fmt.Println(heap)
}
