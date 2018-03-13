package leetcode0152

func maxProduct(nums []int) int {
	res, pos, neg := nums[0], 1, 1
	for _, num := range nums {
		if num == 0 {
			pos, neg = 0, 1
		} else if num > 0 {
			pos *= num
			neg *= num
		} else {
			pos, neg = neg*num, pos*num
		}
		if pos > res {
			res = pos
		}
		if pos <= 0 {
			pos = 1
		}
	}
	return res
}
