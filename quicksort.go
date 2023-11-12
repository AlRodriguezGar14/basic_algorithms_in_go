package main

import "fmt"

func qs(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partitioner(arr, lo, hi)

	qs(arr, pivotIdx+1, hi)
	qs(arr, lo, pivotIdx-1)
}

func partitioner(arr []int, lo int, hi int) int {
	idx := lo - 1
	pivot := arr[hi]

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot
	return idx
}

func testQuickSort() {
	arr := []int{23, 43, 234, 65, 58, 87, 89, 98, 32, 45, 12, 21, 23, 32, 34, 43, 45, 54, 56, 76, 65, 67, 90, 0, 9, 3, 1, 2, 6, 50, 66, 55, 33, 22, 88, 66, 99}

	fmt.Println("Sorted array:")
	fmt.Println(arr)
	qs(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:")
	fmt.Println(arr)

}
