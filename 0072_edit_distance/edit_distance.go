package leetcode0072

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dtab := make([]int, m+1)
	for i := 0; i < m+1; i++ {
		dtab[i] = i
	}

	for j := 1; j < n+1; j++ {
		last := dtab[0]
		dtab[0] = j

		for i := 1; i < m+1; i++ {
			if word1[i-1] == word2[j-1] {
				dtab[i], last = last, dtab[i]
			} else {
				min := last
				if dtab[i] < min {
					min = dtab[i]
				}
				if dtab[i-1] < min {
					min = dtab[i-1]
				}
				dtab[i], last = min+1, dtab[i]
			}
		}
	}
	return dtab[m]
}
