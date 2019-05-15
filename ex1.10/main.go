package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"strings"
	"net/url"
)

const (
	defaultSchema = "http://"
	logfile = "log.txt"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, uri := range os.Args[1:] {
		go fetch(uri, ch)
	}
	for range os.Args[1:] {
		fmt.Printf("%v\n", <-ch)
	}
	fmt.Printf("%0.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(uri string, ch chan<- string) {
	start := time.Now()
	if !strings.HasSuffix(uri, defaultSchema) {
		uri = defaultSchema + uri
	}

	resp, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	f, err := os.Create(url.QueryEscape(uri))
	if err != nil {
		ch <- err.Error()
	}

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()

	if CloseErr := f.Close(); err == nil {
		err = CloseErr
	}

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", uri, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, uri)
}
