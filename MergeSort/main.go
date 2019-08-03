package main

import (
	"fmt"
)

func main() {

	//slice := generateSlice(5)
	var items = []int{2, 7, 4, 8, 1, 5, 3}
	fmt.Println("\n --- unsorted --- \n\n", items)
	fmt.Print("\n--- sorted ---\n\n", mergeSort(items), "\n")
}

func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2
	return merge(mergeSort(slice[:mid]), mergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	count := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			slice[count] = left[i]
			count, i = count+1, i+1
		} else {
			slice[count] = right[j]
			count, j = count+1, j+1

		}
	}

	for i < len(left) {
		slice[count] = left[i]
		count, i = count+1, i+1

	}
	for j < len(right) {
		slice[count] = right[j]
		count, j = count+1, j+1

	}
	return slice
}
