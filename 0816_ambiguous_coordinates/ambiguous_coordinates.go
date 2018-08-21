package leetcode0816

func ambiguousCoordinates(S string) []string {
	S = S[1 : len(S)-1]
	res := []string{}
	for i := 1; i < len(S); i++ {
		p1, p2 := helper(S[:i]), helper(S[i:])
		for _, a := range p1 {
			for _, b := range p2 {
				res = append(res, "("+a+", "+b+")")
			}
		}
	}
	return res
}

func helper(s string) []string {
	n := len(s)
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{s}
	}
	if s[0] == '0' {
		if s[n-1] == '0' {
			return []string{}
		}
		return []string{s[:1] + "." + s[1:]}
	}
	res := []string{s}
	if s[n-1] == '0' {
		return res
	}
	for i := 1; i < n; i++ {
		res = append(res, s[:i]+"."+s[i:])
	}
	return res
}
