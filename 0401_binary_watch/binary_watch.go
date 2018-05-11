package leetcode0401

import "fmt"

func readBinaryWatch(num int) []string {
	hs := []int{8, 4, 2, 1}
	ms := []int{32, 16, 8, 4, 2, 1}
	res := []string{}
	for i := 0; i <= 4 && i <= num; i++ {
		if num-i > 6 {
			continue
		}
		l1, l2 := gen(hs, i), gen(ms, num-i)
		for _, h := range l1 {
			if h >= 12 {
				continue
			}
			for _, m := range l2 {
				if m >= 60 {
					continue
				}
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return res
}

func gen(nums []int, count int) []int {
	res := []int{}
	n := len(nums)
	var helper func(int, int, int)
	helper = func(c, index, sum int) {
		if c == 0 {
			res = append(res, sum)
			return
		}
		for i := index; i < n; i++ {
			helper(c-1, i+1, sum+nums[i])
		}
	}
	helper(count, 0, 0)
	return res
}
