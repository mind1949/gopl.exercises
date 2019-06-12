package counter

import (
	"strings"
	"bufio"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	_w := 0
	for scanner.Scan() {
		_w++
	}
	*w += WordCounter(_w)
	return _w, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	_l := 0
	for scanner.Scan() {
		_l++
	}
	*l += LineCounter(_l)
	return _l, nil
}
