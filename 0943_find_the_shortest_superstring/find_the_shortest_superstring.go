package leetcode0943

import (
	"math"
	"strings"
)

func shortestSuperstring(A []string) string {
	n := len(A)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			graph[i][j] = calc(A[i], A[j])
			graph[j][i] = calc(A[j], A[i])
		}
	}
	nn := 1 << uint(n)
	dp, path := make([][]int, nn), make([][]int, nn)
	for i := 0; i < nn; i++ {
		dp[i] = make([]int, n)
		path[i] = make([]int, n)
	}

	last := -1
	min := math.MaxInt32

	for i := 1; i < nn; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
		for j := 0; j < n; j++ {
			nj := 1 << uint(j)
			if i&nj != 0 {
				prev := i - nj
				if prev == 0 {
					dp[i][j] = len(A[j])
				} else {
					for k := 0; k < n; k++ {
						if dp[prev][k] < math.MaxInt32 && dp[prev][k]+graph[k][j] < dp[i][j] {
							dp[i][j] = dp[prev][k] + graph[k][j]
							path[i][j] = k
						}
					}
				}
			}
			if i == nn-1 && dp[i][j] < min {
				min = dp[i][j]
				last = j
			}
		}
	}

	cur := nn - 1
	stack := []int{}
	for cur > 0 {
		stack = append(stack, last)
		tmp := cur
		cur -= 1 << uint(last)
		last = path[tmp][last]
	}

	bs := strings.Builder{}
	p := len(stack) - 1
	i := stack[p]
	p--
	bs.WriteString(A[i])
	for p >= 0 {
		j := stack[p]
		p--
		bs.WriteString(A[j][len(A[j])-graph[i][j]:])
		i = j
	}
	return bs.String()
}

func calc(a, b string) int {
	m, n := len(a), len(b)
	for i := 1; i < m; i++ {
		if strings.HasPrefix(b, a[i:]) {
			return n - m + i
		}
	}
	return n
}
