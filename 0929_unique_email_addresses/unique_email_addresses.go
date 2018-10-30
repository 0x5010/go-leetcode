package leetcode0929

import "strings"

func numUniqueEmails(emails []string) int {
	m := map[string]struct{}{}
	for _, email := range emails {
		n := len(email)
		bs := strings.Builder{}
		ignore := false
		for i := 0; i < n; i++ {
			if email[i] == '@' {
				bs.WriteString(email[i:])
				break
			}
			if ignore || email[i] == '.' {
				continue
			}
			if email[i] == '+' {
				ignore = true
				continue
			}
			bs.WriteByte(email[i])
		}
		m[bs.String()] = struct{}{}
	}
	return len(m)
}
