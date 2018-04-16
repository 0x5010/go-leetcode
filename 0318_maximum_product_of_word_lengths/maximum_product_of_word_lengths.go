package leetcode0318

func maxProduct(words []string) int {
	n := len(words)

	bits := make([]int, n)
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			bits[i] |= 1 << uint32(word[j]-'a')
		}
	}

	res := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if bits[i]&bits[j] != 0 {
				continue
			}
			tmp := len(words[i]) * len(words[j])
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}
