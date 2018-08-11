package leetcode0777

func canTransform(start string, end string) bool {
	n := len(start)
	if len(end) != n {
		return false
	}
	for p1, p2 := 0, 0; p1 < n && p2 < n; p1, p2 = p1+1, p2+1 {
		for p1 < n && start[p1] == 'X' {
			p1++
		}
		for p2 < n && end[p2] == 'X' {
			p2++
		}
		if p1 == n && p2 == n {
			return true
		}
		if p1 == n || p2 == n ||
			start[p1] != end[p2] ||
			start[p1] == 'L' && p2 > p1 ||
			start[p1] == 'R' && p1 > p2 {
			return false
		}
	}
	return true
}
