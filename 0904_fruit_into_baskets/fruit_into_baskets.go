package leetcode0904

func totalFruit(tree []int) int {
	var a, b, cur, cb int
	res := 0
	for _, c := range tree {
		if c == a || c == b {
			cur++
		} else {
			cur = cb + 1
		}
		if c == b {
			cb++
		} else {
			cb = 1
			a, b = b, c
		}
		if cur > res {
			res = cur
		}
	}
	return res
}
