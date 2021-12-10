package main

import "strings"

type lexer struct {
    input string
    start int
}

func (l *lexer) next() rune {
    l.start += 1
    if l.start >= len(l.input) {
        return -1
    }
    return rune(l.input[l.start])
}

func (l *lexer) back() {
    l.start -= 1
}

func (l *lexer) acceptRun(s string) int {
    count := 0
    for strings.ContainsRune(s, l.next()) {
        count++
    }
    l.back()
    return count
}

func (l *lexer) accept(s string) bool {
    if strings.ContainsRune(s, l.next()) {
        return true
    }
    l.back()
    return false
}

func NewLexer(s string) *lexer {
    return &lexer{
        input: s,
        start: -1,
    }
}

// 基于状态机的解析器
// 首先找到数字的文法表示形式
func isNumber(s string) bool {
    input := strings.TrimSpace(s)
    if input == "" {
        return false
    }
    l := NewLexer(input)
    digits := "0123456789"
    // 吞掉前置的
    l.accept("+-")
    // 吞掉数字部分
    prefixNum := l.acceptRun(digits)
    suffixNum := 0

    // 小数部分搞定
    if l.accept(".") {
        suffixNum = l.acceptRun(digits)
        if prefixNum == 0 && suffixNum == 0 {
            return false
        }
    }

    // 科学计数部分搞定
    if l.accept("Ee") {
        if prefixNum == 0 && suffixNum == 0 {
            return false
        }
        l.accept("+-")
        if l.acceptRun(digits) == 0 {
            return false
        }
    }
    return l.next() == -1
}

