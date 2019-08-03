package main

import "fmt"

func main() {
	haystack := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(9, haystack))
}

func binarySearch(needle int, array []int) (bool, int) {
	low := 0
	high := len(array) - 1

	for low < high {
		median := (low + high) / 2

		if array[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if array[low] != needle || low == len(array) {
		return false, -1
	}
	return true, low
}
