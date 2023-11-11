package main

import "fmt"

func bubble_sort(arr []int) {
	var tmp int
	var end int = len(arr)

	for ; end > 0; end-- {
		for i := 0; i < end-1; i++ {
			if arr[i] > arr[i+1] {
				tmp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp

			}
		}
	}
}

func bubble_sort_v2(arr []int) {
	var tmp int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp

			}
		}
	}
}

func print_bubble_sort() {
	var arr = []int{100, 37, 5, 3, 50, 7, 21, 3, 9, 8, 1}
	bubble_sort_v2(arr)
	fmt.Println("The sorted array is:", arr)

}
