package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrack(n, 0, 0, "", &res)
	return res
}

func backtrack(n int, left int, right int, s string, res *[]string) {
	if right == n {
		*res = append(*res, s)
		return
	}
	if left < n {
		backtrack(n, left+1, right, s+"(", res)
	}
	if right < left {
		backtrack(n, left, right+1, s+")", res)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
