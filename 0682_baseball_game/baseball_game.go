package leetcode0682

import "strconv"

func calPoints(ops []string) int {
	s := []int{}

	for i := range ops {
		switch ops[i] {
		case "+":
			s = append(s, s[len(s)-1]+s[len(s)-2])
		case "D":
			s = append(s, 2*s[len(s)-1])
		case "C":
			s = s[:len(s)-1]
		default:
			point, _ := strconv.Atoi(ops[i])
			s = append(s, point)
		}
	}

	res := 0
	for _, p := range s {
		res += p
	}
	return res
}
