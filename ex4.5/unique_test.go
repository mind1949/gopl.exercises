package unique

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	list := []struct {
		input, want []string
	}{
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "a", "a", "a", "b"}, []string{"a", "b"}},
		{[]string{"a", "a", "b", "b", "b", "c", "c", "c", "c", "d", "d", "d", "d", "d"}, []string{"a", "b", "c", "d"}},
	}

	for _, e := range list {
		got := unique(e.input)
		if !reflect.DeepEqual(got, e.want) {
			t.Errorf("got: %v, want: %v", got, e.want)
		}
	}
}
