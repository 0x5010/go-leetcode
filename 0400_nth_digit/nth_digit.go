package leetcode0400

import (
	"strconv"
)

func findNthDigit(n int) int {
	n--
	first, digits := 1, 1
	for n/9/first/digits >= 1 {
		n -= 9 * first * digits
		digits++
		first *= 10
	}
	return int(strconv.Itoa(first + n/digits)[n%digits] - '0')
}
