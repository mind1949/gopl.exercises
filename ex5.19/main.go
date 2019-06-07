// ex5.19使用panic与recover写一个函数，他没有return语句，但是能够返回一个非零的值．
package main

import (
	"fmt"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()

	panic("try panic and recover")
}
