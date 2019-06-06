package ints

import (
	"testing"
)

func TestMax(t *testing.T) {
	datas := []struct {
		input []int
		want  int
	}{
		{input: []int{}, want: 0},
		{input: []int{1, 2, 3, 4, 5}, want: 5},
	}

	for _, data := range datas {
		got := max(data.input...)
		if got != data.want {
			t.Errorf("want: %v, got: %v", data.want, got)
		}
	}
}

func TestMin(t *testing.T) {
	datas := []struct {
		input []int
		want  int
	}{
		{input: []int{}, want: 0},
		{input: []int{1, 2, 3, 4, 5, 0}, want: 0},
	}

	for _, data := range datas {
		got := min(data.input...)
		if got != data.want {
			t.Errorf("want: %v, got: %v", data.want, got)
		}
	}
}
