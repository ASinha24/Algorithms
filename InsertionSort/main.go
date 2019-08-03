package main

import (
	"fmt"
)

func main() {
	var nums = []int{2, 7, 4, 9, 3, 1}
	insertionSort(nums)
	fmt.Println(nums)
}

func insertionSort(nums []int) {
	var n = len(nums)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
			j = j - 1
		}
	}
}
