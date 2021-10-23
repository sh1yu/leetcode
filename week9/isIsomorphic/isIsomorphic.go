package main

import "fmt"

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}

	transbook := make(map[byte]byte)
	alreadySet := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if c, ok := transbook[s[i]]; ok {
			if c != t[i] {
				return false
			}
		} else {
			if alreadySet[t[i]] {
				return false
			}
			transbook[s[i]] = t[i]
			alreadySet[t[i]] = true
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("badc", "baba"))
}
