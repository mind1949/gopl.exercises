package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		os.Exit(1)
	}

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "os.Open(%s) %v\n", fileName, err)
			continue
		}

		if hasRepeatedLines(file) {
			fmt.Printf("%s has repeated\n", fileName)
		}
	}
}

func hasRepeatedLines(file *os.File) (ret bool) {
	scanner := bufio.NewScanner(file)
	lines := map[string]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		if lines[line] {
			ret = true
			break
		}
		lines[line] = true
	}

	return ret
}
