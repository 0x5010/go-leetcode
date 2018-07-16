package leetcode0583

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)

	dp := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		pre := dp[0]
		for j := 0; j < n2; j++ {
			tmp := dp[j+1]
			if word1[i] == word2[j] {
				dp[j+1] = pre + 1
			} else {
				dp[j+1] = max(dp[j+1], dp[j])
			}
			pre = tmp
		}
	}

	return n1 + n2 - 2*dp[n2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
