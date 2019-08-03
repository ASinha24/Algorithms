package main

import "fmt"

func linearSearchInt(numbers []int, key int) int {
	for i, v := range numbers {
		if key == v {
			return i
		}
	}
	return -1
}

func linearSearchString(str []string, skey string) int {
	for i, v := range str {
		if skey == v {
			return i
		}
	}
	return -1
}

func main() {
	numbers := []int{34, 78, 95, 78, 82, 46, 58, 45, 94, 34, 86, 99, 251, 320}
	str := []string{"Abyssinian", "American", "Cornish", "Devon", "Oriental"}
	indexString := linearSearchString(str, "Cornish")
	indexInt := linearSearchInt(numbers, 45)
	if indexString == -1 {
		fmt.Println("string is not present")
	} else {
		fmt.Println("string is present at ", indexString)
	}
	if indexInt == -1 {
		fmt.Println("integer is not present")
	} else {
		fmt.Println("integer is present at ", indexInt)
	}
}
