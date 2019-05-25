// 练习4.1　编写一个函数，　用于统计SHA256散列中不同的位数
package shadiff

import (
	"crypto/sha256"
)

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func bitDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += popCount(b[i])
		case i >= len(b):
			count += popCount(a[i])
		default:
			count += popCount(a[i] ^ b[i])
		}
	}
	return count
}

func ShaBitDiff(a, b []byte) int {
	c1 := sha256.Sum256(a)
	c2 := sha256.Sum256(b)
	return bitDiff(c1[:], c2[:])
}
