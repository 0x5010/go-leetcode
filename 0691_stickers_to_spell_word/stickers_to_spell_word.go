package leetcode0691

import "math"

func minStickers(stickers []string, target string) int {
	n := len(stickers)

	ss := make([][26]int, n)
	for i, s := range stickers {
		for _, c := range s {
			ss[i][c-'a']++
		}
	}

	dp := map[string]int{}
	dp[""] = 0

	var helper func(string) int
	helper = func(target string) int {
		if minimum, ok := dp[target]; ok {
			return minimum
		}

		tar := [26]int{}
		for _, c := range target {
			tar[c-'a']++
		}

		res := math.MaxInt32

		for _, s := range ss {
			if s[target[0]-'a'] == 0 {
				continue
			}
			t := make([]byte, 0, len(target))
			for i := 0; i < 26; i++ {
				for j := tar[i] - s[i]; 0 < j; j-- {
					t = append(t, byte('a'+i))
				}
			}
			tmp := helper(string(t))
			if tmp != -1 {
				res = min(res, tmp+1)
			}
		}
		if res == math.MaxInt32 {
			res = -1
		}
		dp[target] = res
		return res
	}
	helper(target)

	return dp[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
