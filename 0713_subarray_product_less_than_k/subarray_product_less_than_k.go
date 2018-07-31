package leetcode0713

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k < 2 {
		return 0
	}
	res := 0
	pro := 1
	l := 0
	for r, num := range nums {
		pro *= num
		for pro >= k {
			pro /= nums[l]
			l++
		}
		res += r - l + 1
	}
	return res
}
