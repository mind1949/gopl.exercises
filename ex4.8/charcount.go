package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := map[string]int{}

	file, err := os.Open("data.text")
	check(err)
	reader := bufio.NewReader(file)

	for {
		r, n, err := reader.ReadRune()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			fmt.Println("invalid rune: ", string(r))
			counts["invalid"]++
		}
		if unicode.IsLetter(r) {
			fmt.Println("Letter rune: ", string(r))
			counts["letter"]++
		}
	}

	for k, v := range counts {
		fmt.Printf("%q's count is %d\n", k, v)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
