package main

import (
	"testing"
)

var slice = []string{"a", "b", "c"}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Concat(slice)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join(slice)
	}
}
