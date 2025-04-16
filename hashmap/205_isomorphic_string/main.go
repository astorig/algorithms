package main

import "fmt"

func main() {
	s := "add"
	t := "egg"

	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
	sLen := len(s)
	mapS := [256]byte{}
	mapT := [256]byte{}

	for i := 0; i < sLen; i++ {

		fmt.Println(mapS)
		fmt.Println(mapT)
		if mapS[s[i]] != mapT[t[i]] {
			return false
		}

		mapS[s[i]] = byte(i + 1)
		mapT[t[i]] = byte(i + 1)
	}

	return true
}
