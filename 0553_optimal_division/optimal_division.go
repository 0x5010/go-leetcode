package leetcode0553

import "strconv"

func optimalDivision(nums []int) string {
	n := len(nums)
	if n == 0 {
		return ""
	}
	res := strconv.Itoa(nums[0])
	if n == 1 {
		return res
	}
	if n == 2 {
		return res + "/" + strconv.Itoa(nums[1])
	}
	res += "/(" + strconv.Itoa(nums[1])
	for i := 2; i < n; i++ {
		res += "/" + strconv.Itoa(nums[i])
	}
	return res + ")"
}
