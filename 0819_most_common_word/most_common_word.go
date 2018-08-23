package leetcode0819

func mostCommonWord(paragraph string, banned []string) string {
	count := map[string]int{}
	m := map[string]struct{}{}
	for _, b := range banned {
		m[b] = struct{}{}
	}
	bs := []byte{}
	for i := 0; i < len(paragraph); i++ {
		b := paragraph[i]
		if b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' {
			if b <= 'Z' {
				b += 32
			}
			bs = append(bs, b)
			continue
		}
		if len(bs) == 0 {
			continue
		}
		word := string(bs)
		if _, ok := m[word]; !ok {
			count[word]++
		}
		bs = bs[:0]
	}
	if len(bs) != 0 {
		word := string(bs)
		if _, ok := m[word]; !ok {
			count[word]++
		}
		bs = bs[:0]
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
