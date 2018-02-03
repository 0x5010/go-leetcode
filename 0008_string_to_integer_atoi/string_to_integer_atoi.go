package leetcode0008

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	neg := false
	if str[0] == '-' {
		neg = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	res := 0
	for _, c := range str {
		if c < '0' || c > '9' {
			break
		}
		val := int(c - '0')
		if neg {
			if (-1 * res) < (math.MinInt32+val)/10 {
				return math.MinInt32
			}
		} else if res > (math.MaxInt32-val)/10 {
			return math.MaxInt32
		}
		res = 10*res + val

	}
	if neg {
		return -1 * res
	}
	return res
}
