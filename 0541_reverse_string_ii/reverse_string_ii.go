package leetcode0541

func reverseStr(s string, k int) string {
	bs := []byte(s)
	n := len(s)
	for i := 0; i < n; i += 2 * k {
		j := min(i+k, n)
		reverse(bs[i:j])
	}
	return string(bs)
}

func reverse(bs []byte) {
	i, j := 0, len(bs)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
