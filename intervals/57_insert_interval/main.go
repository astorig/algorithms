package main

import (
	"fmt"
	"sort"
)

func main() {

	data := [][]int{
		{1, 3},
		{6, 9},
	}

	newInterval := []int{2, 5}

	fmt.Println(insert(data, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= merged[len(merged)-1][1] {
			if intervals[i][1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = intervals[i][1]
			} else {
				continue
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}
