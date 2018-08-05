package leetcode0739

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	s := make([]int, n)
	top := -1
	res := make([]int, n)
	for i := 0; i < n; i++ {
		for top > -1 && temperatures[i] > temperatures[s[top]] {
			res[s[top]] = i - s[top]
			top--
		}
		top++
		s[top] = i
	}
	return res
}
