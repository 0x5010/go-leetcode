package leetcode0726

import (
	"sort"
	"strconv"
	"strings"
)

func countOfAtoms(formula string) string {
	n := len(formula)
	m := map[string]int{}
	s := []map[string]int{}
	var val int
	getNum := func(i int) (int, int) {
		val := 0
		for i < n && isDigit(formula[i]) {
			val = val*10 + int(formula[i]-'0')
			i++
		}
		if val == 0 {
			val = 1
		}
		return i, val
	}

	for i := 0; i < n; {
		b := formula[i]
		i++
		if b == '(' {
			s = append(s, m)
			m = map[string]int{}
		} else if b == ')' {
			i, val = getNum(i)
			if len(s) != 0 {
				tmp := m
				m = s[len(s)-1]
				s = s[:len(s)-1]
				for k, v := range tmp {
					m[k] += v * val
				}
			}
		} else {
			start := i - 1
			for i < n && idLower(formula[i]) {
				i++
			}
			atom := formula[start:i]
			i, val = getNum(i)
			m[atom] += val
		}
	}
	bs := strings.Builder{}
	atoms := make([]string, 0, len(m))
	for atom := range m {
		atoms = append(atoms, atom)
	}
	sort.Strings(atoms)
	for _, atom := range atoms {
		bs.WriteString(atom)
		val = m[atom]
		if val > 1 {
			bs.WriteString(strconv.Itoa(val))
		}
	}
	return bs.String()
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func idLower(b byte) bool {
	return 'a' <= b && b <= 'z'
}
