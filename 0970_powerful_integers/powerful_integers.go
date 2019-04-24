package leetcode0970

func powerfulIntegers(x int, y int, bound int) []int {
	s := map[int]struct{}{}
	for i := 1; i <= bound; i *= x {
		for j := 1; i+j <= bound; j *= y {
			s[i+j] = struct{}{}
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	l := make([]int, 0, len(s))
	for k := range s {
		l = append(l, k)
	}
	return l
}
