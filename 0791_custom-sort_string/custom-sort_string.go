package leetcode0791

import "bytes"

func customSortString(S string, T string) string {
	count := [26]int{}
	for i := 0; i < len(T); i++ {
		count[T[i]-'a']++
	}
	bs := bytes.Buffer{}
	for i := 0; i < len(S); i++ {
		for count[S[i]-'a'] > 0 {
			bs.WriteByte(S[i])
			count[S[i]-'a']--
		}
	}
	for b := byte('a'); b <= 'z'; b++ {
		for count[b-'a'] > 0 {
			bs.WriteByte(b)
			count[b-'a']--
		}
	}
	return bs.String()
}
