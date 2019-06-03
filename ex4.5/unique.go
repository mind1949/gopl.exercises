// 编写一个就地处理函数，用于去除[]string slice中相邻的重复字符串元素．
package unique

func unique(s []string) []string {
	c := 0
	for _, e := range s {
		if s[c] == e {
			continue
		}
		c++
		s[c] = e
	}
	return s[:c+1]
}
