package leetcode0862

func shortestSubarray(A []int, K int) int {
	n := len(A)
	sums := make([]int, n+1)
	for i, a := range A {
		if a == K {
			return 1
		}
		sums[i+1] = sums[i] + a
	}
	deque := make([]int, n+2)
	res := n + 1
	for i, l, r := 0, 0, -1; i < n+1; i++ {
		for r-l >= 0 && sums[i]-sums[deque[l]] >= K {
			res = min(res, i-deque[l])
			l++
		}
		for r-l >= 0 && sums[i] <= sums[deque[r]] {
			r--
		}
		r++
		deque[r] = i
	}
	if res == n+1 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
