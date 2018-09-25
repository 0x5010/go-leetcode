package leetcode0020

func isValid(s string) bool {
	n := len(s)

	match := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}

	for i := 0; i < n; i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			m := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if m != match[s[i]] {
				return false
			}
		}
	}
	return len(stack) == 0
}
