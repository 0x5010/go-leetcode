package leetcode0925

func isLongPressedName(name string, typed string) bool {
	m, n, i := len(name), len(typed), 0
	for j := 0; j < n; j++ {
		if i < m && name[i] == typed[j] {
			i++
		} else if j == 0 || typed[j] != typed[j-1] {
			return false
		}
	}
	return i == m
}
