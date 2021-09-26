package main

import "github.com/sh1yu/assertion"

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += int((num >> i) & 1)
	}
	return count
}

func main() {
	assertion.Equals(3, hammingWeight(0b00000000000000000000000000001011))
}
