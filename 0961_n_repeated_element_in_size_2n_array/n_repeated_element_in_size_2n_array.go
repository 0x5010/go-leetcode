package leetcode0961

func repeatedNTimes(A []int) int {
	m := map[int]struct{}{}
	for _, a := range A {
		if _, ok := m[a]; ok {
			return a
		}
		m[a] = struct{}{}
	}
	return -1
}
