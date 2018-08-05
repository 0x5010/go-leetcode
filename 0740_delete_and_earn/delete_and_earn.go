package leetcode0740

func deleteAndEarn(nums []int) int {
	sum := make([]int, 10001)
	for _, num := range nums {
		sum[num] += num
	}
	a, b := 0, 0
	res := 0
	for _, v := range sum {
		if a > b+v {
			res = a
		} else {
			res = b + v
		}
		a, b = res, a
	}
	return res
}
