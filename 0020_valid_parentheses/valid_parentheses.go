package leetcode0020

var match = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

type stack []rune

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) pop() (rune, bool) {
	if len(*s) > 0 {
		res := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return res, true
	}
	return 0, false
}

func isValid(s string) bool {
	st := new(stack)
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			st.push(r)
		case ')', ']', '}':
			if m, ok := st.pop(); !ok || m != match[r] {
				return false
			}
		}
	}

	if len(*st) > 0 {
		return false
	}
	return true
}
