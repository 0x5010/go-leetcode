package leetcode0227

func calculate(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	stack := make([]int, 0, n)
	num := 0
	var sign byte = '+'

	for i := 0; i < n; i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}
		if (!(c >= '0' && c <= '9') && c != ' ') || i == n-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -1*num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			sign = c
			num = 0
		}
	}
	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}
