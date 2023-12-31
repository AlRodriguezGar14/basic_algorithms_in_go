package main

import "fmt"

// This algo only works with sorted arrays
func binary_search_old(arr []int, lo int, hi int, n int) bool {

	for lo < hi {

		var mid int = lo + (hi-lo)/2 // This is basically lo + half the distance to hi
		v := arr[mid]

		if v == n {
			return true
		} else if v < n {
			lo = mid + 1
		} else {
			hi = mid
		}
		binary_search_old(arr, lo, hi, n)
	}
	return false
}

func binary_search(arr []int, n int) (bool, int) {

	var lo int = 0
	var hi int = len(arr)

	for lo < hi {

		mid := lo + (hi-lo)/2
		v := arr[mid]

		if v == n {
			return true, mid
		} else if v < n {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return false, 0
}

func print_binary_search() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	found, index := binary_search(arr, 12)
	fmt.Printf("Is the number in the array? %t\n", found)
	if found {
		fmt.Printf("At the index: %d", index)
	}

}
