package leetcode0819

import "bytes"

func mostCommonWord(paragraph string, banned []string) string {
	count := map[string]int{}
	m := map[string]struct{}{}
	for _, b := range banned {
		m[b] = struct{}{}
	}
	bs := bytes.Buffer{}
	for i := 0; i < len(paragraph); i++ {
		b := paragraph[i]
		if b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' {
			if b <= 'Z' {
				b += 32
			}
			bs.WriteByte(b)
			continue
		}
		if bs.Len() == 0 {
			continue
		}
		word := bs.String()
		if _, ok := m[word]; !ok {
			count[word]++
		}
		bs.Reset()
	}
	if bs.Len() != 0 {
		word := bs.String()
		if _, ok := m[word]; !ok {
			count[word]++
		}
		bs.Reset()
	}
	max := 0
	res := ""
	for word, c := range count {
		if c > max {
			max = c
			res = word
		}
	}
	return res
}
