package leetcode0795

func numSubarrayBoundedMax(A []int, L int, R int) int {
	heads, tails := 0, 0
	res := 0
	for _, a := range A {
		if a < L {
			tails++
			res += heads
		} else if a <= R {
			heads += tails + 1
			tails = 0
			res += heads
		} else {
			heads, tails = 0, 0
		}
	}
	return res
}
