package leetcode0224

func calculate(s string) int {
	res := 0
	stack := make([]int, 0, len(s))
	sign := 1

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			res += sign * num
			i--
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			stack = append(stack, res, sign)
			res = 0
			sign = 1
		case ')':
			sign = stack[len(stack)-1]
			tmp := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			res = sign*res + tmp
		}
	}
	return res
}
