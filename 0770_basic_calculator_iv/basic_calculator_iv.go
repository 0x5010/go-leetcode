package leetcode0770

import (
	"sort"
	"strconv"
	"strings"
)

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	eemap := make(map[string]int, len(evalvars))
	for i := range evalvars {
		eemap[evalvars[i]] = evalints[i]
	}

	expression = strings.Replace(expression, " ", "", -1)

	numbers := parse(expression, eemap)

	return format(numbers)
}

func parse(expression string, m map[string]int) nums {
	n1, n2 := nums{num{c: 0}}, nums{num{c: 0}}
	opt1, opt2 := byte('+'), byte('+')

	for {
		var n3 nums
		var i int
		if expression[0] == '(' {
			i = indexOfCounterParentheses(expression)
			n3 = parse(expression[1:i], m)
			i++
		} else {
			i = indexOfNextOpt(expression)
			n3 = creatNums(expression[:i], m)
		}

		if opt2 == '*' {
			n2 = operate(opt2, n2, n3)
		} else {
			n1 = operate(opt1, n1, n2)
			n2 = n3
			opt1 = opt2
		}

		if i == len(expression) {
			break
		}

		opt2 = expression[i]
		expression = expression[i+1:]
	}

	return operate(opt1, n1, n2)
}

func creatNums(exp string, m map[string]int) nums {
	constant, ok := m[exp]
	if ok {
		return nums{num{c: constant}}
	}

	constant, err := strconv.Atoi(exp)
	if err != nil {
		return nums{num{vars: []string{exp}, c: 1}}
	}

	return nums{num{c: constant}}
}

func operate(opt byte, a, b nums) nums {
	var res nums
	switch opt {
	case '+':
		res = add(a, b)
	case '-':
		res = minus(a, b)
	case '*':
		res = mult(a, b)
	}
	return res
}

func add(a, b nums) nums {
	return append(a, b...)
}

func minus(a, b nums) nums {
	for i := range b {
		b[i].c *= -1
	}
	return append(a, b...)
}

func mult(a, b nums) nums {
	res := make(nums, 0, len(a)*len(b))
	for i := range a {
		for j := range b {
			vars := make([]string, 0, len(a[i].vars)+len(b[j].vars))
			vars = append(vars, a[i].vars...)
			vars = append(vars, b[j].vars...)
			res = append(res, num{vars: vars, c: a[i].c * b[j].c})
		}
	}
	return res
}

func indexOfCounterParentheses(expression string) int {
	i := 1
	count := 1
	for ; i < len(expression); i++ {
		switch expression[i] {
		case '(':
			count++
		case ')':
			count--
		}
		if count == 0 {
			break
		}
	}
	return i
}

func indexOfNextOpt(expression string) int {
	var i int
	for i = 1; i < len(expression); i++ {
		if expression[i] == '+' ||
			expression[i] == '-' ||
			expression[i] == '*' {
			break
		}
	}
	return i
}

func format(numbers nums) []string {
	numbers = update(numbers)
	res := make([]string, 0, len(numbers))

	for _, n := range numbers {
		if n.c == 0 {
			continue
		}
		temp := strconv.Itoa(n.c)
		if n.key != "" {
			temp = temp + "*" + n.key
		}
		res = append(res, temp)
	}

	return res
}

type nums []num

type num struct {
	vars []string
	key  string
	c    int
}

func update(ns nums) nums {
	for i := range ns {
		sort.Strings(ns[i].vars)
		ns[i].key = strings.Join(ns[i].vars, "*")
	}

	sort.Slice(ns, func(i int, j int) bool {
		leni := len(ns[i].vars)
		lenj := len(ns[j].vars)
		if leni != lenj {
			return leni > lenj
		}

		for k := 0; k < leni; k++ {
			if ns[i].vars[k] == ns[j].vars[k] {
				continue
			}
			return ns[i].vars[k] < ns[j].vars[k]
		}
		return false
	})

	res := nums{ns[0]}
	i, j := 0, 1
	for j < len(ns) {
		if res[i].key == ns[j].key {
			res[i].c += ns[j].c
		} else {
			res = append(res, ns[j])
			i++
		}
		j++
	}

	return res
}
