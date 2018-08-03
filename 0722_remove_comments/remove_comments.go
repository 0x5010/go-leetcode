package leetcode0722

import "bytes"

func removeComments(source []string) []string {
	res := []string{}
	bs := bytes.Buffer{}
	mode := false
	for _, s := range source {
		n := len(s)
		for i := 0; i < n; i++ {
			if mode {
				if s[i] == '*' && i < n-1 && s[i+1] == '/' {
					mode = false
					i++
				}
			} else {
				if s[i] == '/' && i < n-1 {
					if s[i+1] == '/' {
						break
					} else if s[i+1] == '*' {
						mode = true
						i++
						continue
					}
				}
				bs.WriteByte(s[i])
			}
		}
		if mode == false && bs.Len() != 0 {
			res = append(res, bs.String())
			bs = bytes.Buffer{}
		}
	}
	return res
}
