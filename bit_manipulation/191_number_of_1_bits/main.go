package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 11

	fmt.Println(hammingWeight(n))
}

func hammingWeight(n int) int {
	br := strconv.FormatInt(int64(n), 2)

	return strings.Count(br, "1")
}
