package leetcode0738

import "strconv"

func monotoneIncreasingDigits(N int) int {
	bs := []byte(strconv.Itoa(N))
	n := len(bs)
	mark := n
	for i := n - 1; i > 0; i-- {
		if bs[i] < bs[i-1] {
			mark = i
			bs[i-1] = bs[i-1] - 1
		}
	}
	for i := mark; i < n; i++ {
		bs[i] = '9'
	}
	res, _ := strconv.Atoi(string(bs))
	return res
}
