package leetcode0482

func licenseKeyFormatting(S string, K int) string {
	bs := []byte{}
	for i := len(S) - 1; i >= 0; i-- {
		b := S[i]
		if b == '-' {
			continue
		}
		if len(bs)%(K+1) == K {
			bs = append(bs, '-')
		}
		if b >= 'a' {
			b -= 32
		}
		bs = append(bs, b)
	}
	i, j := 0, len(bs)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
