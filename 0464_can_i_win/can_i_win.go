package leetcode0464

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	sum := (1 + maxChoosableInteger) * maxChoosableInteger / 2
	if desiredTotal > sum {
		return false
	}

	bit := make([]int, maxChoosableInteger+1)
	for i := maxChoosableInteger; i > 0; i-- {
		bit[i] = 1 << uint8(i)
	}

	dp := make(map[int]bool, maxChoosableInteger*maxChoosableInteger)

	var dfs func(int, int) bool
	dfs = func(usedBit, remains int) bool {
		if res, ok := dp[usedBit]; ok {
			return res
		}
		if remains <= maxChoosableInteger {
			for i := maxChoosableInteger; i >= remains; i-- {
				if usedBit&bit[i] == 0 {
					dp[usedBit] = true
					return true
				}
			}
		}
		for i := maxChoosableInteger; i > 0; i-- {
			if usedBit&bit[i] > 0 {
				continue
			}
			tmp := usedBit
			usedBit |= bit[i]
			lose := dfs(usedBit, remains-i)
			usedBit = tmp

			if !lose {
				dp[usedBit] = true
				return true
			}
		}
		dp[usedBit] = false
		return false
	}
	return dfs(0, desiredTotal)
}
