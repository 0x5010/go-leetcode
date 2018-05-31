package leetcode0442

func findDuplicates(nums []int) []int {
	res := []int{}
	for _, num := range nums {
		n := num
		if n < 0 {
			n *= -1
		}
		if nums[n-1] < 0 {
			res = append(res, n)
		} else {
			nums[n-1] *= -1
		}
	}
	return res
}
