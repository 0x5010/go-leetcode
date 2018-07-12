package leetcode0557

import "strings"

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	for i, s := range ss {
		ss[i] = reverse(s)
	}
	return strings.Join(ss, " ")
}

func reverse(s string) string {
	i, j := 0, len(s)-1
	bs := []byte(s)
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
