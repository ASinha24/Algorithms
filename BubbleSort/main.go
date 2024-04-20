package main

import (
	"fmt"
)

func main() {
	var items = []int{2, 7, 4, 1, 5, 3}
	bubbleSort(items)
	fmt.Println(items)
}

func bubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		sorted = true
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				sorted = false
			}
		}
	}
}
