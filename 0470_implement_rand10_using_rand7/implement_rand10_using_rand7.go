package leetcode0470

func rand10() int {
	x := 40
	for x >= 40 {
		x = 7*(rand7()-1) + (rand7() - 1)
	}
	return x%10 + 1
}
