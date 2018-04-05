package leetcode0299

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	nums := [10]int{}
	for i := 0; i < len(secret); i++ {
		s, g := int(secret[i]-'0'), int(guess[i]-'0')
		if s == g {
			bulls++
		} else {

			if nums[s] < 0 {
				cows++
			}
			if nums[g] > 0 {
				cows++
			}
			nums[s]++
			nums[g]--
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
