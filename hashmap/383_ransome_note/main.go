package main

func main() {
	ransomeNote := "avcv"
	magazine := "bbbcv"

	canConstruct(ransomeNote, magazine)
}

func canConstruct(ransomNote string, magazine string) bool {
	magazineLetters := make(map[rune]int)

	for _, letter := range magazine {
		magazineLetters[letter]++
	}

	for _, letter := range ransomNote {
		if magazineLetters[letter] == 0 {
			return false
		}

		magazineLetters[letter]--
	}

	return true
}
