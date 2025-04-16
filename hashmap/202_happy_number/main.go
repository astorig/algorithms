package main

func main() {
	n := 19

	isHappy(n)
}

func isHappy(n int) bool {
	numbers := map[int]bool{}

	for n != 1 && !numbers[n] {
		numbers[n] = true

		sum := 0

		for n > 0 {
			rest := n % 10
			sum += rest * rest
			n = n / 10
		}

		n = sum
	}

	return n == 1
}
