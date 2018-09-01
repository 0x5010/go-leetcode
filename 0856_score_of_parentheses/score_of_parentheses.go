package leetcode0856

func scoreOfParentheses(S string) int {
	n := len(S)
	f := 1
	res := 0
	for i := 0; i < n; i++ {
		if S[i] == '(' {
			f *= 2
		} else {
			f /= 2
		}
		if S[i] == '(' && S[i+1] == ')' {
			f /= 2
			res += f
			i++
		}
	}
	return res
}
