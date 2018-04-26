package leetcode0345

func reverseVowels(s string) string {
	n := len(s)
	bs, i, j := []byte(s), 0, n-1
	for {
		for i < n-1 && !isVowel(bs[i]) {
			i++
		}
		for j > 0 && !isVowel(bs[j]) {
			j--
		}
		if i >= j {
			break
		}
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
