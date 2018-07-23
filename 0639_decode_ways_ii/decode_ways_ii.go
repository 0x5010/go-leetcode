package leetcode0639

func numDecodings(s string) int {
	m := 1000000007
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, one(s[0])
	for i := 2; i <= n; i++ {
		w1, w2 := one(s[i-1]), two(s[i-2], s[i-1])
		dp[i] = dp[i-1]*w1 + dp[i-2]*w2
		if dp[i] == 0 {
			return 0
		}
		dp[i] %= m
	}
	return dp[n]
}

func one(a byte) int {
	if a == '*' {
		return 9
	}
	if a == '0' {
		return 0
	}
	return 1
}

func two(a, b byte) int {
	if a == '*' {
		if b == '*' {
			return 15
		}
		if '0' <= b && b <= '6' {
			return 2
		}
		return 1
	}
	if a == '1' {
		if b == '*' {
			return 9
		}
		return 1
	}
	if a == '2' {
		if b == '*' {
			return 6
		}
		if b == '7' || b == '8' || b == '9' {
			return 0
		}
		return 1
	}
	return 0
}
