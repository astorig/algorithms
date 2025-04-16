package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7

	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	sum := 0

	minLength := math.MaxInt32

	left := 0
	right := 0
	for right < len(nums) {
		sum = sum + nums[right]

		for sum >= target {
			minLength1 := (right + 1) - left

			if minLength > minLength1 {
				minLength = minLength1
			}
			sum -= nums[left]
			left++
		}
		right++
	}

	if minLength == math.MaxInt32 {
		return 0
	}

	return minLength
}
