package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	s := []int{5, 4, 3, 2, 1}
	RotateInts(s)
	want := []int{4, 3, 2, 1, 5}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("go %v want %v", s, want)
	}
}
