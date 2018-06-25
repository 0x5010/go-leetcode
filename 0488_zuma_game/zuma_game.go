package leetcode0488

const maxCount = 6

func findMinStep(board string, hand string) int {
	ball := make([]int, 26)
	for i := 0; i < len(hand); i++ {
		ball[hand[i]-'A']++
	}

	var dfs func(string) int
	dfs = func(s string) int {
		s = removeConsecutive(s)
		if s == "#" {
			return 0
		}
		count, need := maxCount, 0
		for i, j := 0, 0; j < len(s); j++ {
			if s[i] == s[j] {
				continue
			}
			need = i + 3 - j
			if ball[s[i]-'A'] >= need {
				ball[s[i]-'A'] -= need
				count = min(count, need+dfs(s[:i]+s[j:]))
				ball[s[i]-'A'] += need
			}
			i = j
		}
		return count
	}
	res := dfs(board + "#")
	if res == maxCount {
		return -1
	}
	return res
}

func removeConsecutive(board string) string {
	for i, j := 0, 0; j < len(board); j++ {
		if board[i] == board[j] {
			continue
		}
		if i+3 <= j {
			return removeConsecutive(board[:i] + board[j:])
		}
		i = j
	}
	return board
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
