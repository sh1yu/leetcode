package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	token := strings.Split(strings.TrimSpace(s), " ")
	token = reverse(token)
	return strings.Join(token, " ")
}

func reverse(strs []string) []string {
	res := make([]string, 0, len(strs))
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] == "" {
			continue
		}
		res = append(res, strs[i])
	}
	return res
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("a good   example"))
}
