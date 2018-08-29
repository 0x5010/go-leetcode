package leetcode0842

import (
	"math"
	"strconv"
)

func splitIntoFibonacci(S string) []int {
	n := len(S)
	res := []int{}
	var recur func(int) bool
	recur = func(index int) bool {
		if index == n && len(res) >= 3 {
			return true
		}
		for i := index; i < n; i++ {
			if S[index] == '0' && i > index {
				break
			}
			num, err := strconv.Atoi(S[index : i+1])
			if err != nil || num > math.MaxInt32 {
				break
			}
			size := len(res)
			if size >= 2 && num > res[size-1]+res[size-2] {
				break
			}
			if size <= 1 || num == res[size-1]+res[size-2] {
				res = append(res, num)
				if recur(i + 1) {
					return true
				}
				res = res[:size]
			}
		}
		return false
	}
	recur(0)
	return res
}
