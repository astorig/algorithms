package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	number, _ := strconv.ParseInt(a, 2, 64)

	number2, _ := strconv.ParseInt(b, 2, 64)

	return strconv.FormatInt(number+number2, 2)

}
