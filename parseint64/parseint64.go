package main

import (
	"errors"
	"fmt"
	"math"
)

func parse(str string) (int64, error) {
	runeStr := []rune(str)
	var isNeg bool
	var result int64
	if len(runeStr) == 0 {
		return -1, errors.New("invalid")
	}
	if runeStr[0] == '-' {
		isNeg = true
		runeStr = runeStr[1:]
	}
	if len(runeStr) == 0 {
		return -1, errors.New("invalid")
	}

	littleMaxInt64 := int64(math.MaxInt64) / 10
	littleMinInt64 := int64(uint64(-math.MinInt64) / 10)
	for _, c := range runeStr {
		if c < '0' || c > '9' {
			return -1, errors.New("invalid")
		}
		cur := int64(c - '0')
		if !isNeg && (result > littleMaxInt64 || result == littleMaxInt64 && cur > math.MaxInt64%10) {
			return -1, errors.New("overflow")
		}
		if isNeg && (result > littleMinInt64 || result == littleMinInt64 && cur > int64(uint64(-math.MinInt64)%10)) {
			return -1, errors.New("overflow")
		}
		result = result*10 + cur
	}
	if isNeg {
		result = -result
	}
	return result, nil
}

func main() {
	// 12A - err
	// 12.3 - err
	// -12
	fmt.Println(parse("233"))
	fmt.Println(parse("-233"))
	fmt.Println(parse("0"))
	fmt.Println(parse("-0"))
	fmt.Println(parse("-"))
	fmt.Println(parse("99999999999999999999999999999999999999999999999999999999999999999999999999999"))
	fmt.Println(parse("-99999999999999999999999999999999999999999999999999999999999999999999999999999"))
	//fmt.Println(math.MaxInt64)
	fmt.Println(parse("9223372036854775807"))
	fmt.Println("====")
	fmt.Println(parse("-9223372036854775808"))
	fmt.Println(parse("9223372036854775808"))
	fmt.Println(parse("-9223372036854775809"))
}
