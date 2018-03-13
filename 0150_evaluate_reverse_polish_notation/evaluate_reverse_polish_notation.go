package leetcode0150

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens)/2+1)
	for _, s := range tokens {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			n := len(stack)
			a, b := stack[n-2], stack[n-1]
			stack[n-2] = calculate(a, b, s)
			stack = stack[:n-1]
		} else {
			tmp, _ := strconv.Atoi(s)
			stack = append(stack, tmp)
		}
	}
	return stack[0]
}

func calculate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}
