package leetcode0831

import "strings"

func maskPII(S string) string {
	if strings.ContainsRune(S, '@') {
		return maskEmail(strings.ToLower(S))
	}
	return maskPhone(S)
}

func maskEmail(s string) string {
	ss := strings.Split(s, "@")
	ss[0] = ss[0][0:1] + "*****" + ss[0][len(ss[0])-1:]
	return strings.Join(ss, "@")
}

func maskPhone(s string) string {
	bs := []byte{}
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			bs = append(bs, s[i])
		}
	}
	n := len(bs)
	s = string(bs[n-4:])

	if n == 10 {
		return "***-***-" + s
	}
	return "+" + strings.Repeat("*", n-10) + "-***-***-" + s
}
