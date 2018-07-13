package leetcode0575

func distributeCandies(candies []int) int {
	n := len(candies)
	m := make(map[int]struct{}, n)
	for _, candy := range candies {
		m[candy] = struct{}{}
	}
	return min(len(m), n/2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
