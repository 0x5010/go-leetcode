package leetcode0482

import (
	"bytes"
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
	b := bytes.Buffer{}
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		}
		if b.Len()%(K+1) == K {
			b.WriteByte('-')
		}
		b.WriteByte(S[i])
	}

	return strings.ToUpper(reverse(b.Bytes()))
}

func reverse(bs []byte) string {
	i, j := 0, len(bs)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
