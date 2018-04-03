package leetcode0274

func hIndex(citations []int) int {
	n := len(citations)
	counts := make([]int, n+1)
	for _, c := range citations {
		if c > n {
			counts[n]++
		} else {
			counts[c]++
		}
	}
	count := 0
	for i := n; i >= 0; i-- {
		count += counts[i]
		if count >= i {
			return i
		}
	}
	return 0
}
