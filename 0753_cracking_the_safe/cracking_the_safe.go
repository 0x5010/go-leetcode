package leetcode0753

func crackSafe(n int, k int) string {
	target := make([]byte, n-1)
	for i := 0; i < n-1; i++ {
		target[i] = '0'
	}

	pref := map[string]int{}

	for {
		cur := string(target[len(target)-(n-1):])

		if r, ok := pref[cur]; !ok {
			pref[cur] = k - 1
		} else if r < 0 {
			return string(target)
		}

		target = append(target, '0'+byte(pref[cur]))
		pref[cur]--
	}
}
