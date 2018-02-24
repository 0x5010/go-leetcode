package leetcode0076

func minWindow(s string, t string) string {
	sLen, tLen := len(s), len(t)

	need := make([]int, 256)
	for i := range t {
		need[t[i]]++
	}

	has := make([]int, 256)

	minStart, minEnd, count, min := 0, 0, 0, sLen+1

	for start, end := 0, 0; end < sLen; end++ {
		if need[s[end]] == 0 {
			continue
		}

		if has[s[end]] < need[s[end]] {
			count++
		}
		has[s[end]]++

		if count == tLen {
			for need[s[start]] == 0 || has[s[start]] > need[s[start]] {
				if has[s[start]] > need[s[start]] {
					has[s[start]]--
				}
				start++
			}

			tmp := end - start + 1
			if tmp < min {
				min = tmp
				minStart = start
				minEnd = end + 1
			}
		}

	}

	if count < tLen {
		return ""
	}
	return s[minStart:minEnd]

}
