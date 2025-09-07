package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}

	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	uniqueMap := map[int]struct{}{}

	for _, num := range nums {
		if _, ok := uniqueMap[num]; ok {
			return true
		}

		uniqueMap[num] = struct{}{}
	}

	return false
}
