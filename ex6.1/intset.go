//ex6.1 实现下列方法：
// func (*Inset) Len() int //返回元素个数
// func (*Inset) Remove(x int) // 从集合去除元素x
// func (*Inset) Clear() // 删除所有元素
// func (*Inset) Copy() *Inset // 返回集合的副本
package intset

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint64(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint64(x%64)
	if word > len(s.words) || s.words[word]&(1<<uint(bit)) == 0 {
		return
	}
	s.words[word] &^= 1 << uint(bit)
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	new := new(IntSet)
	new.words = make([]uint64, len(s.words))
	copy(new.words, s.words) // 若用new.words = s.words就会让两者引用同一个数组了，　这就不算是副本了
	return new
}
