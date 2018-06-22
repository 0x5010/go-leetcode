package leetcode0474

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		numZeros, numOnes := 0, 0
		for i := 0; i < len(str); i++ {
			if str[i] == '0' {
				numZeros++
			} else {
				numOnes++
			}
		}

		for i := m; i >= numZeros; i-- {
			for j := n; j >= numOnes; j-- {
				tmp := dp[i-numZeros][j-numOnes] + 1
				if tmp > dp[i][j] {
					dp[i][j] = tmp
				}
			}
		}
	}
	return dp[m][n]
}
