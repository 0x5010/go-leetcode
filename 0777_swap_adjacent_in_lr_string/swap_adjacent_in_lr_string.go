package leetcode0777

func canTransform(start string, end string) bool {
	n := len(start)
	if len(end) != n {
		return false
	}
	for i, j := 0, 0; i < n && j < n; i, j = i+1, j+1 {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}
		if i == n && j == n {
			return true
		}
		if i == n || j == n ||
			start[i] != end[j] ||
			start[i] == 'L' && j > i ||
			start[i] == 'R' && i > j {
			return false
		}
	}
	return true
}
