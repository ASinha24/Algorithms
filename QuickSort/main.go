package main

import (
	"fmt"
)

func main() {
	var items = []int{2, 5, 8, 3, 3, 6, 6, 9, 10, 4}
	quickSort(items, 0, len(items)-1)
	fmt.Println(items)
}

func quickSort(items []int, startIndex, endIndex int) {
	if startIndex < endIndex {
		partitionIndex := partition(items, startIndex, endIndex)
		quickSort(items, startIndex, partitionIndex-1)
		quickSort(items, partitionIndex+1, endIndex)
	}
}

func partition(items []int, startIndex, endIndex int) int {
	pivot := items[endIndex]
	pIndex := startIndex
	for j := startIndex; j <= endIndex-1; j++ {
		if items[j] < pivot {
			items[pIndex], items[j] = items[j], items[pIndex]
			pIndex++
		}
	}
	items[pIndex], items[endIndex] = items[endIndex], items[pIndex]
	return pIndex
}
