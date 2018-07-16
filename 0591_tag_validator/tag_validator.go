package leetcode0591

import "strings"

func isValid(code string) bool {
	validTagName := func(i, j int) bool {
		if i > j || i == j || i+9 < j {
			return false
		}
		for k := i; k < j; k++ {
			if code[k] < 'A' || 'Z' < code[k] {
				return false
			}
		}
		return true
	}

	n := len(code)
	stack := []string{}
	i := 0
	for i < n {
		if i > 0 && len(stack) == 0 {
			return false
		}

		switch {
		case i+9 < n && code[i:i+9] == "<![CDATA[":
			i += 9
			j := i + strings.Index(code[i:], "]]>")
			if i > j {
				return false
			}
			i = j + 3
		case i+2 < n && code[i:i+2] == "</":
			i += 2
			j := i + strings.Index(code[i:], ">")
			if !validTagName(i, j) {
				return false
			}
			tagName := code[i:j]
			i = j + 1
			n := len(stack)
			if n == 0 || stack[n-1] != tagName {
				return false
			}
			stack = stack[:n-1]
		case code[i] == '<':
			i++
			j := i + strings.Index(code[i:], ">")
			if !validTagName(i, j) {
				return false
			}
			tagName := code[i:j]
			i = j + 1
			stack = append(stack, tagName)
		default:
			i++
		}
	}

	return len(stack) == 0
}
