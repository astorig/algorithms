package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}

	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	sumMap := map[int]int{}

	for i, num := range nums {
		if _, ok := sumMap[target-num]; ok {
			return []int{sumMap[target-num], i}
		}

		sumMap[num] = i
	}

	return []int{}
}
