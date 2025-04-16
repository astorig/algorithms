package main

import (
	"fmt"
)

const caseDelta = 'a' - 'A'

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str))
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if skipChar(s[start]) {
			start++
			continue
		}

		if skipChar(s[end]) {
			end--
			continue
		}

		if toLower(s[start]) != toLower(s[end]) {
			return false
		}

		start++
		end--
	}

	return true
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + caseDelta
	}

	if char >= 'a' && char <= 'z' {

	}

	return char
}

func skipChar(char byte) bool {
	return (char < 'a' || char > 'z') &&
		(char < 'A' || char > 'Z') &&
		(char < '0' || char > '9')
}
