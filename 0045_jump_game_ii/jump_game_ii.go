package leetcode0045

func jump(nums []int) int {
	start, end, step, n := 0, 0, 0, len(nums)
	for end < n-1 {
		step++
		maxEnd := end + 1
		for i := start; i < end+1; i++ {
			tmp := i + nums[i]
			if tmp >= n-1 {
				return step
			}
			if tmp > maxEnd {
				maxEnd = tmp
			}
		}
		start, end = end+1, maxEnd
	}
	return step
}
