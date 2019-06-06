package join

import (
	"bytes"
)

func Join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	var buf bytes.Buffer
	for _, s := range strs {
		buf.WriteString(s)
		buf.WriteString(sep)
	}
	buf.Truncate(buf.Len() - len(sep))
	return buf.String()
}
