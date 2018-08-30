package leetcode0847

import "math"

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	maskSize := 1 << uint(n)

	queue := [][2]int{}
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, maskSize)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
		mask := 1 << uint(i)
		dp[i][mask] = 0
		queue = append(queue, [2]int{i, mask})
	}

	for len(queue) != 0 {
		s, m := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, next := range graph[s] {
			nextMask := m | 1<<uint(next)
			if dp[next][nextMask] > dp[s][m]+1 {
				dp[next][nextMask] = dp[s][m] + 1
				queue = append(queue, [2]int{next, nextMask})
			}
		}
	}

	res := math.MaxInt32
	traversalMask := maskSize - 1
	for i := 0; i < n; i++ {
		res = min(res, dp[i][traversalMask])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
