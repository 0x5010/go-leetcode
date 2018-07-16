package leetcode0592

import (
	"strconv"
	"strings"
)

func fractionAddition(expression string) string {
	fractions := split(expression)
	res := newFraction(fractions[0])
	for i := 1; i < len(fractions); i++ {
		res = add(res, newFraction(fractions[i]))
	}
	return res.String()
}

func split(s string) []string {
	n := len(s)
	i, j := 0, 1
	res := []string{}
	for j < n {
		if s[j] == '+' {
			res = append(res, s[i:j])
			i = j + 1
		} else if s[j] == '-' {
			res = append(res, s[i:j])
			i = j
		}
		j++
	}
	res = append(res, s[i:j])
	return res
}

type fraction struct {
	numer, denom int
}

func newFraction(s string) *fraction {
	ss := strings.Split(s, "/")
	n, _ := strconv.Atoi(ss[0])
	d, _ := strconv.Atoi(ss[1])
	return &fraction{n, d}
}

func add(a, b *fraction) *fraction {
	n := a.numer*b.denom + b.numer*a.denom
	d := a.denom * b.denom
	x := gcd(abs(n), d)
	return &fraction{n / x, d / x}
}

func (f *fraction) String() string {
	return strconv.Itoa(f.numer) + "/" + strconv.Itoa(f.denom)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
