package leetcode0640

import (
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	es := strings.Split(equation, "=")
	la, lb := evaluate(es[0])
	ra, rb := evaluate(es[1])
	a, b := la-ra, rb-lb

	if a == 0 {
		if b == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}
	return "x=" + strconv.Itoa(b/a)
}

func evaluate(exp string) (int, int) {
	n := len(exp)
	i, j := 0, 1
	a, b := 0, 0
	parse := func(token string) {
		if token[len(token)-1] == 'x' {
			if token == "x" {
				a++
			} else if token == "-x" {
				a--
			} else {
				t, _ := strconv.Atoi(token[:len(token)-1])
				a += t
			}
		} else {
			t, _ := strconv.Atoi(token)
			b += t
		}
	}

	for j < n {
		if exp[j] == '+' {
			parse(exp[i:j])
			i = j + 1
		} else if exp[j] == '-' {
			parse(exp[i:j])
			i = j
		}
		j++
	}
	parse(exp[i:j])
	return a, b
}
