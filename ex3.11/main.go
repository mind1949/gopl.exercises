// 增强comma函数的功能，让其能够正确处理浮点数，以及带有可选正负号的数字
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func comma(s string) string {
	var buf bytes.Buffer
	r := reverse([]rune(s))
	for count, i := 0, 0; i < len(r); i++ {
		buf.WriteRune(r[i])
		if unicode.IsDigit(r[i]) {
			count++
			if count%3 == 0 && i != len(r)-1 {
				buf.WriteString(",")
			}
		}
	}

	r = []rune(buf.String())
	return string(reverse(r))
}

func reverse(r []rune) []rune {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func main() {
	var s string
	fmt.Print("enter an number(int or float): \n\t")
	fmt.Scanf("%s", &s)
	fmt.Println(comma(s))
}
