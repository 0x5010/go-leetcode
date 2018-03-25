package leetcode0214

func shortestPalindrome(s string) string {
	n := len(s)
	rs := reverse(s)
	mirror := s + "#" + rs
	next := make([]int, 2*n+1)
	for i := 1; i < 2*n+1; i++ {
		j := next[i-1]
		for j > 0 && mirror[i] != mirror[j] {
			j = next[j-1]
		}
		if mirror[i] == mirror[j] {
			j++
		}
		next[i] = j

	}
	return rs[:n-next[2*n]] + s
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
