// 编写一个函数判断两个字符串是否同文异构，也就是，他们都含有相同的字符但序列不同
package anagram

func isAnagram(a, b string) bool {
	aFreq := make(map[rune]int)
	for _, c := range a {
		aFreq[c]++
	}

	bFreq := make(map[rune]int)
	for _, c := range b {
		bFreq[c]++
	}

	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}

	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}
	return true
}
