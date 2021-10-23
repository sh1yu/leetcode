package main

import (
	"bytes"
	"github.com/sh1yu/assertion"
)

// word := a-zA-Z0-9
// split

func newSplit(str string) []string {

	result := make([]string, 0)
	var buf bytes.Buffer
	for _, s := range str {
		if isAlphaDigit(s) {
			buf.WriteRune(s)
		} else {
			tmp := buf.String()
			if len(tmp) > 0 {
				result = append(result, tmp)
			}
			buf = bytes.Buffer{}
		}
	}
	tmp := buf.String()
	if len(tmp) > 0 {
		result = append(result, tmp)
	}
	return result
}

func isAlphaDigit(b int32) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func main() {
	assertion.Equals(newSplit("11 231 A%%S 3"), []string{"11", "231", "A", "S", "3"})
	assertion.Equals(newSplit("this is a test."), []string{"this", "is", "a", "test"})
	assertion.Equals(newSplit(""), []string{})
	assertion.Equals(newSplit("     "), []string{})
	assertion.Equals(newSplit("33.. ..  33"), []string{"33", "33"})
	assertion.Equals(newSplit("test"), []string{"test"})
	assertion.Equals(newSplit("test^^^"), []string{"test"})
	assertion.Equals(newSplit("^^^test"), []string{"test"})
}
