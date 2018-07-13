package leetcode0567

func checkInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 > n2 {
		return false
	}

	count := 0
	for i := 0; i < n1; i++ {
		count += 1 << (s1[i] - 'a')
		count -= 1 << (s2[i] - 'a')
	}

	if count == 0 {
		return true
	}

	for i := n1; i < n2; i++ {
		count += (1 << (s2[i-n1] - 'a')) - (1 << (s2[i] - 'a'))
		if count == 0 {
			return true
		}
	}

	return false
}
