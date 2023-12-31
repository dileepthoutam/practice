package main

import "fmt"

func SubArrayWithGivenSum(arr []int, target int) (int, int) {
	start := 0
	end := 0
	sum := 0

	for end < len(arr) {
		if sum < target {
			sum += arr[end]
			end = end + 1
		} else if sum > target {
			sum = sum - arr[start]
			start = start + 1
		}

		if sum == target {
			return start, end - 1
		}
	}

	return -1, -1
}

func main() {

	arr := []int{1, 3, 2, 5, 1, 1, 2, 3}
	target := 5

	start, end := SubArrayWithGivenSum(arr, target)

	fmt.Println(start, end)

}
