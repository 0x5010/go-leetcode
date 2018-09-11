package leetcode0880

func decodeAtIndex(S string, K int) string {
	var i, n int
	for i = 0; n < K; i++ {
		if S[i] <= '9' {
			n *= int(S[i] - '0')
		} else {
			n++
		}
	}
	for {
		i--
		if S[i] <= '9' {
			n /= int(S[i] - '0')
			K %= n
		} else {
			if K == 0 || K == n {
				return string(S[i])
			}
			n--
		}
	}
}
