package main

import "github.com/sh1yu/assertion"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func main() {
	assertion.Equals(true, isPowerOfTwo(1))
	assertion.Equals(true, isPowerOfTwo(16))
	assertion.Equals(false, isPowerOfTwo(3))
}
