package leetcode0636

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	s := []int{}
	prev := 0
	for _, log := range logs {
		parts := strings.Split(log, ":")
		i, _ := strconv.Atoi(parts[0])
		t, _ := strconv.Atoi(parts[2])
		if len(s) != 0 {
			res[s[len(s)-1]] += t - prev
		}
		prev = t
		if parts[1] == "start" {
			s = append(s, i)
		} else {
			res[s[len(s)-1]]++
			s = s[:len(s)-1]
			prev++
		}
	}
	return res
}
