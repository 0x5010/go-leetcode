package leetcode0006

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := strings.Builder{}

	p := numRows*2 - 2

	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	for i := 1; i <= numRows-2; i++ {
		res.WriteByte(s[i])

		for j := p; j-i < len(s); j += p {
			res.WriteByte(s[j-i])
			if i+j < len(s) {
				res.WriteByte(s[i+j])
			}
		}
	}

	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}
	return res.String()
}
