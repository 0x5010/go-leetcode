package leetcode0209

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	res, i, j, sum := n+1, 0, 0, 0
	for j < n {
		sum += nums[j]
		j++

		for sum >= s {
			sum -= nums[i]
			i++
			if res > j-i+1 {
				res = j - i + 1
			}
		}
	}
	return res % (n + 1)
}
