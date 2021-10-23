package main

import "fmt"

func numDecodings(s string) int {

	a := 0
	b := 1
	c := 0
	for i := 1; i <= len(s); i++ {
		if s[i-1] > '0' {
			c = b
			if i > 1 && s[i-2] > '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
				c = a + b
			}
		} else {
			if i > 1 && s[i-2] > '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
				c = a
			} else {
				c = 0
			}
		}
		a, b = b, c
	}
	return c
}

func main() {
	fmt.Println(numDecodings("12"))
}
