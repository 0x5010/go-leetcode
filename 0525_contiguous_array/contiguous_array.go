package leetcode0525

func findMaxLength(nums []int) int {
	m := map[int]int{}
	m[0] = -1

	zeros, ones, res := 0, 0, 0

	for i, num := range nums {
		if num == 0 {
			zeros++
		} else {
			ones++
		}

		if j, ok := m[zeros-ones]; ok {
			res = max(res, i-j)
		} else {
			m[zeros-ones] = i
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
