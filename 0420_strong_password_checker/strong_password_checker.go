package leetcode0420

func strongPasswordChecker(s string) int {
	n := len(s)
	l, u, d, missingType := 1, 1, 1, 3
	for i := 0; i < n; i++ {
		if 0 < l && 'a' <= s[i] && s[i] <= 'z' {
			l = 0
			missingType--
		}
		if 0 < u && 'A' <= s[i] && s[i] <= 'Z' {
			u = 0
			missingType--
		}
		if 0 < d && '0' <= s[i] && s[i] <= '9' {
			d = 0
			missingType--
		}

		if missingType == 0 {
			break
		}
	}

	var replace, ones, twos int

	for p := 0; p+2 < len(s); p++ {
		if s[p] != s[p+1] || s[p+1] != s[p+2] {
			continue
		}
		repeatingLen := 2
		for p+2 < len(s) && s[p] == s[p+2] {
			repeatingLen++
			p++
		}

		replace += repeatingLen / 3
		if repeatingLen%3 == 0 {
			ones++
		} else if repeatingLen%3 == 1 {
			twos++
		}
	}

	if len(s) < 6 {
		return max(missingType, 6-len(s))
	}

	if len(s) <= 20 {
		return max(missingType, replace)
	}
	delete := len(s) - 20

	replace -= min(delete, ones)
	replace -= min(max(delete-ones, 0), twos*2) / 2
	replace -= max(delete-ones-2*twos, 0) / 3

	return delete + max(missingType, replace)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
