package main

import "fmt"

func main() {
	s := "pwwkew"

	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[rune]bool)
	maxLength, start := 0, 0

	for end, char := range s {
		for charSet[char] {
			delete(charSet, rune(s[start]))
			start++
		}

		charSet[char] = true
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength

}
