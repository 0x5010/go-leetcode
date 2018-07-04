package leetcode0504

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	res, sign := "", ""
	if num < 0 {
		sign = "-"
		num *= -1
	}
	for num > 0 {
		res = strconv.Itoa(num%7) + res
		num /= 7
	}
	return sign + res
}
