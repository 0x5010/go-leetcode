package leetcode0049

import (
	"bytes"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	r := make(map[string][]string)
	for _, str := range strs {
		count := make([]int, 26)
		for _, b := range []byte(str) {
			count[b-'a']++
		}

		b := bytes.Buffer{}
		for _, i := range count {
			b.WriteByte('#')
			b.WriteByte(byte(i) + '0')
		}
		key := b.String()
		r[key] = append(r[key], str)
	}
	res := make([][]string, len(r))
	i := 0
	for _, l := range r {
		res[i] = l
		i++
	}
	return res
}
