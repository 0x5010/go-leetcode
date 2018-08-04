package leetcode0730

func countPalindromicSubsequences(S string) int {
	mod := 1000000007
	n := len(S)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	before := make([][4]int, n)
	after := make([][4]int, n)

	indexOfBefore := [4]int{n, n, n, n}
	indexOfAfter := [4]int{-1, -1, -1, -1}

	for i := 0; i < n; i++ {
		j := n - i - 1
		indexOfBefore[S[i]-'a'] = i
		indexOfAfter[S[j]-'a'] = j
		before[i] = indexOfBefore
		after[j] = indexOfAfter
	}

	var recur func(int, int) int
	recur = func(l, r int) int {
		if l > r {
			return 0
		}

		if dp[l][r] != 0 {
			return dp[l][r]
		}

		count := 0
		for k := 0; k < 4; k++ {
			i, j := after[l][k], before[r][k]
			if l <= i && j <= r {
				if i == j {
					count++
				} else if i < j {
					count += recur(i+1, j-1) + 2
				}
			}
		}
		count %= mod
		dp[l][r] = count
		return count
	}
	return recur(0, n-1)
}
