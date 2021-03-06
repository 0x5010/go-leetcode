package leetcode0139

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]struct{}, len(wordDict))
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				if _, ok := wordSet[s[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[n]
}
