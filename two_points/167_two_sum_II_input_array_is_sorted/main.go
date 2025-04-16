package main

import "fmt"

func main() {

	numbers := []int{-1, 0}
	target := -1

	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return []int{-1, -1}
}
