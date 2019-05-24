//编写一个非递归的comma函数，运用bytes.buffer, 而不是简单的字符串拼接
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	endIndex := len(s) - 1
	for i, v := range s {
		buf.WriteRune(v)
		if (endIndex-i) % 3 == 0 && i != endIndex  {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

func main() {
	var s string
	fmt.Print("enter an integer: ")
	fmt.Scanf("%s", &s)
	fmt.Println(comma(s))
}
