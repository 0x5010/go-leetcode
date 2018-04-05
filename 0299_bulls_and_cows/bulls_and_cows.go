package leetcode0299

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	nums := [10]int{}
	for i := 0; i < len(secret); i++ {
		s, g := secret[i], guess[i]
		if s == g {
			bulls++
		} else {
			si, gi := int(s)-'0', int(g)-'0'
			if nums[si] < 0 {
				cows++
			}
			if nums[gi] > 0 {
				cows++
			}
			nums[si]++
			nums[gi]--
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
