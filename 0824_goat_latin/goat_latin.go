package leetcode0824

import "strings"

func toGoatLatin(S string) string {
	n := len(S)
	if n == 0 {
		return ""
	}
	m := map[byte]struct{}{}
	for _, vowel := range []byte("aeiouAEIOU") {
		m[vowel] = struct{}{}
	}
	bs := strings.Builder{}
	index := 1
	for _, word := range strings.Fields(S) {
		if index != 1 {
			bs.WriteByte(' ')
		}
		f := word[0]
		if _, ok := m[f]; ok {
			bs.WriteString(word)
		} else {
			bs.WriteString(word[1:])
			bs.WriteByte(f)
		}
		bs.WriteString("ma")
		for i := 0; i < index; i++ {
			bs.WriteByte('a')
		}
		index++
	}
	return bs.String()
}
