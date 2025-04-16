package main

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3

	removeElement(nums, val)
}

func removeElement(nums []int, val int) int {
	result := 0

	for _, num := range nums {
		if num != val {
			nums[result] = num
			result++
		}
	}

	return result
}
