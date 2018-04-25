package leetcode0344

func reverseString(s string) string {
	bs, i, j := []byte(s), 0, len(s)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
