package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const defaultschema = "http://"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "missing url!\n")
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		if !strings.HasSuffix(url, defaultschema) {
			url = defaultschema + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		_, err = io.Copy(os.Stdout, strings.NewReader(resp.Status+"\n"))
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %v: %v\n", url, err)
			os.Exit(1)
		}
	}
}
