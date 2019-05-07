package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(Join(os.Args[1:]))
	fmt.Println(Concat(os.Args[1:]))
}

func Concat(slice []string) string {
	var s, sep string
	for _, arg := range slice {
		s += sep + arg
		sep = " "
	}
	return s
}

func Join(slice []string) string {
	return strings.Join(slice, " ")
}
