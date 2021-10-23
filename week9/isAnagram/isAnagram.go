package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var hash [26]int
	for _, c := range s {
		hash[c-'a']++
	}
	for _, c := range t {
		hash[c-'a']--
	}
	for _, count := range hash {
		if count > 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
