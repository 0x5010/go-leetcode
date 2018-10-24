package leetcode0926

func minFlipsMonoIncr(S string) int {
	n := len(S)
	if n == 0 {
		return 0
	}
	countOne := 0
	res := 0
	for i := 0; i < n; i++ {
		if S[i] == '0' {
			if countOne == 0 {
				continue
			} else {
				res++
			}
		} else {
			countOne++
		}
		if res > countOne {
			res = countOne
		}
	}
	return res
}
