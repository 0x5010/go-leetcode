package leetcode0736

import "strconv"

func evaluate(expression string) int {
	return eval(expression, nil)
}

func eval(exp string, e *env) int {
	if exp[0] != '(' {
		num, err := strconv.Atoi(exp)
		if err != nil {
			return e.get(exp)
		}
		return num
	}

	exp = exp[1 : len(exp)-1]
	identifier, exp := parse(exp)
	switch identifier {
	case "add":
		a, b := parse(exp)
		return eval(a, e) + eval(b, e)
	case "mult":
		a, b := parse(exp)
		return eval(a, e) * eval(b, e)
	case "let":
		e = newEnv(e)
		var key, val string
		for {
			key, exp = parse(exp)
			if exp == "" {
				break
			}
			val, exp = parse(exp)
			e.set(key, eval(val, e))
		}
		return eval(key, e)
	}
	return 0
}

func parse(exp string) (string, string) {
	n := len(exp)
	i := 0
	if exp[0] == '(' {
		l := 0
		for i < n {
			if exp[i] == '(' {
				l++
			} else if exp[i] == ')' {
				l--
			}
			if l == 0 {
				break
			}
			i++
		}
	} else {
		for i+1 < n {
			if exp[i+1] == ' ' {
				break
			}
			i++
		}
	}

	a, b := exp[:i+1], ""
	if i+1 != n {
		b = exp[i+2:]
	}

	return a, b
}

type env struct {
	parent *env
	m      map[string]int
}

func newEnv(parent *env) *env {
	return &env{
		parent: parent,
		m:      map[string]int{},
	}
}

func (e *env) set(k string, v int) {
	e.m[k] = v
}

func (e *env) get(k string) int {
	if v, ok := e.m[k]; ok {
		return v
	}
	if e.parent != nil {
		return e.parent.get(k)
	}
	return 0
}
