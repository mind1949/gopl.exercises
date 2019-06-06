package join

import (
	"testing"
)

func TestJoin(t *testing.T) {
	datas := []struct {
		strs []string
		sep  string
		want string
	}{
		{strs: []string{}, sep: ",", want: ""},
		{strs: []string{"1", "2", "3", "4"}, sep: ",", want: "1,2,3,4"},
		{strs: []string{"a", "b"}, sep: "|", want: "a|b"},
	}

	for _, data := range datas {
		got := Join(data.sep, data.strs...)
		if got != data.want {
			t.Errorf("got: %v, want: %v\n", got, data.want)
		}
	}
}
