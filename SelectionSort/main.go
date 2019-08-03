package main

import (
	"fmt"
)

func main() {
	var nums = []int{9, 7, 8, 3, 4, 2, 5}
	selectionSort(nums)
	fmt.Println(nums)
}

func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := 0; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
			nums[j], nums[minIndex] = nums[minIndex], nums[j]
		}
	}
}
