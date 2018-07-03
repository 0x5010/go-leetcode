package leetcode0503

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	s := make([]int, 0, n)
	for i := 0; i < n*2; i++ {
		for len(s) != 0 && nums[s[len(s)-1]] < nums[i%n] {
			res[s[len(s)-1]] = nums[i%n]
			s = s[:len(s)-1]
		}
		s = append(s, i%n)
	}
	return res
}
