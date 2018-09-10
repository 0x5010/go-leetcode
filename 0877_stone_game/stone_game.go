package leetcode0877

func stoneGame(piles []int) bool {
	return true
}

// func stoneGame(piles []int) bool {
// 	n := len(piles)
// 	dp := make([][]int, n)
// 	for i, p := range piles {
// 		dp[i] = make([]int, n)
// 		dp[i][i] = p
// 	}
// 	for d := 1; d < n; d++ {
// 		for i := 0; i < n-d; i++ {
// 			dp[i][i+d] = max(piles[i]-dp[i+1][i+d], piles[i+d]-dp[i][i+d-1])
// 		}
// 	}
// 	return dp[0][n-1] > 0
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
