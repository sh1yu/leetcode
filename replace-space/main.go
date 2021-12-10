package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {
    s := "We are happy."
    fmt.Println(replaceSpace(s))

    reg := regexp.MustCompile("^[+-]?((([0-9]*\\.)?[0-9]+([Ee][+-]?[0-9]+)?)|([0-9]+(\\.[0-9]*)?([Ee][+-]?[0-9]+)?))$")
    fmt.Println(reg.MatchString("005047e+6"))
}

func replaceSpace(s string) string {
    var buf bytes.Buffer
    for _, c := range s {
        if c == ' ' {
            buf.WriteString("%20")
        } else {
            buf.WriteByte(byte(c))
        }
    }
    return buf.String()
}


