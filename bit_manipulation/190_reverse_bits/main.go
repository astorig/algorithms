package main

func main() {
	var n uint32

	n = 0b00000010100101000001111010011100

	reverseBits(n)
}

func reverseBits(num uint32) uint32 {
	var result uint32

	for i := 0; i < 32; i++ {
		result = result << 1
		result += num & 1
		num = num >> 1
	}

	return result
}
