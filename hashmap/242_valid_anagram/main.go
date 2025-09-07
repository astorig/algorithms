package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	symbolMap := map[rune]int{}

	if len(s) != len(t) {
		return false
	}

	for _, num := range s {
		symbolMap[num]++
	}

	for _, num := range t {
		if _, ok := symbolMap[num]; !ok {
			return false
		}

		symbolMap[num]--

		if symbolMap[num] < 0 {
			return false
		}
	}

	return true
}
