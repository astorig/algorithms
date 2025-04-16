package main

import "fmt"

func main() {

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1

	amountofWater := 0

	for left < right {
		amountofWater = max(amountofWater, min(height[left], height[right])*(right-left))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return amountofWater
}
