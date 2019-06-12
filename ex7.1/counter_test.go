package counter

import "testing"

func TestWordCounter(t *testing.T) {
	p := []byte("我　的　世 界")
	w := new(WordCounter)
	w.Write(p)
	if int(*w) != 4 {
		t.Logf("expect: %d, got: %d", 4, int(*w))
		t.Fail()
	}
}

func TestLineCounter(t *testing.T) {
	p := []byte(`i love word
你好啊
of course i still love you`)
	l := new(LineCounter)
	l.Write(p)
	if int(*l) != 3 {
		t.Logf("expect: %d, got: %d", 3, int(*l))
		t.Fail()
	}
}
