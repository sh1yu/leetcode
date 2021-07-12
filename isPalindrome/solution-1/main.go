package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("a."))
}

//思路：双指针，从前往后遍历
//时间复杂度：O(n)
//空间复杂度：O(1)
func isPalindrome(s string) bool {
	rs := []rune(strings.ToLower(s))
	start := 0
	end := len(rs) - 1
	for start < end {
		if !unicode.IsLetter(rs[start]) && !unicode.IsNumber(rs[start]) {
			start++
			continue
		}
		if !unicode.IsLetter(rs[end]) && !unicode.IsNumber(rs[end]) {
			end--
			continue
		}
		if rs[start] != rs[end] {
			return false
		}

		start++
		end--
	}
	return true
}
