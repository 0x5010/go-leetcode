package leetcode0006

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	bs := []byte{}
	p := numRows*2 - 2
	for i := 0; i < len(s); i += p {
		bs = append(bs, s[i])
	}
	for i := 1; i <= numRows-2; i++ {
		bs = append(bs, s[i])

		for j := p; j-i < len(s); j += p {
			bs = append(bs, s[j-i])
			if i+j < len(s) {
				bs = append(bs, s[i+j])
			}
		}
	}
	for i := numRows - 1; i < len(s); i += p {
		bs = append(bs, s[i])
	}
	return string(bs)
}
