package main

import (
	"fmt"
)

func reverseStr(s string, k int) string {
	res := make([]rune, 0)
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if i%(2*k) == 0 {

			if len(runes)-i >= k {
				res = append(res, reverse(runes[i:i+k])...)
			} else {
				res = append(res, reverse(runes[i:])...)
			}

			if len(runes)-i >= 2*k {
				res = append(res, runes[i+k:i+2*k]...)
			} else if len(runes)-i >= k {
				res = append(res, runes[i+k:]...)
			}
		}
	}

	return string(res)
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}
