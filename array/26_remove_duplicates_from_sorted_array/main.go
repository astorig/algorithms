package main

func main() {
	nums := []int{1, 1, 2}

	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {

	index := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
