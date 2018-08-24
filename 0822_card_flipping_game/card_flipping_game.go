package leetcode0822

func flipgame(fronts []int, backs []int) int {
	m := map[int]struct{}{}
	for i, f := range fronts {
		if f == backs[i] {
			m[f] = struct{}{}
		}
	}
	res := 2001
	for i, f := range fronts {
		if _, ok := m[f]; !ok && res > f {
			res = f
		}
		if _, ok := m[backs[i]]; !ok && res > backs[i] {
			res = backs[i]
		}
	}
	if res == 2001 {
		return 0
	}
	return res
}
