package main

import "fmt"

func main() {
	s := "()[]"

	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		if match, ok := pairs[char]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == match {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
