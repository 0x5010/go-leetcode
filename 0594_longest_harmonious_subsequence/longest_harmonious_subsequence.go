package leetcode0594

func findLHS(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}

	res := 0
	for _, num := range nums {
		if v, ok := m[num+1]; ok {
			res = max(res, m[num]+v)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
