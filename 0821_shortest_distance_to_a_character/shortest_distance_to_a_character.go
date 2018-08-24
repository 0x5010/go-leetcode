package leetcode0821

func shortestToChar(S string, C byte) []int {
	n := len(S)
	pos := -n
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if S[i] == C {
			pos = i
		}
		res[i] = i - pos
	}
	for i := n - 1; i >= 0; i-- {
		if S[i] == C {
			pos = i
		}
		res[i] = min(res[i], abs(i-pos))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
