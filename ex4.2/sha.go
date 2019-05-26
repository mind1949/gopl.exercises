package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

type shaSum func(b []byte) []byte

var width = flag.Int("width", 256, "输入哈希值大小")

func main() {
	flag.Parse()

	sum256 := func(b []byte) []byte {
		h := sha256.Sum256(b)
		return h[:]
	}
	sum512 := func(b []byte) []byte {
		h := sha512.Sum512(b)
		return h[:]
	}

	functions := map[int]shaSum{256: sum256, 512: sum512}
	function := functions[*width]
	if function == nil {
		log.Fatal("不支持此哈希值大小")
	}

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		b := scanner.Bytes()
		h := function(b)
		fmt.Printf("%x\n", h)
	}
}
