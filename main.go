package main

import (
	"fmt"
)

func LinearSearch(arr []int, target int) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}

	return false
}

func BinarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	result := -1

	for start <= end {
		mid := start + (end - start) / 2

		if target == arr[mid] {
			result = mid
			break
		} else if target < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return result
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	target := 100

	fmt.Println(LinearSearch(arr, target))
	
}
