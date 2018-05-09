package leetcode0394

import "strconv"

func decodeString(s string) string {
	n := len(s)

	i := 0
	for i < n && (s[i] < '0' || s[i] > '9') {
		i++
	}
	if i == n {
		return s
	}

	j := i + 1

	for s[j] != '[' {
		j++
	}

	k := j
	count := 1
	for count > 0 {
		k++
		if s[k] == '[' {
			count++
		} else if s[k] == ']' {
			count--
		}
	}

	num, _ := strconv.Atoi(s[i:j])
	return s[:i] + times(decodeString(s[j+1:k]), num) + decodeString(s[k+1:])
}

func times(s string, n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}
