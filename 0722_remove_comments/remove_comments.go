package leetcode0722

func removeComments(source []string) []string {
	res := []string{}
	bs := []byte{}
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
				bs = append(bs, s[i])
			}
		}
		if mode == false && len(bs) != 0 {
			res = append(res, string(bs))
			bs = bs[:0]
		}
	}
	return res
}
