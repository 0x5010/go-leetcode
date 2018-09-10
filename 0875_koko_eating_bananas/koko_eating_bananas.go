package leetcode0875

func minEatingSpeed(piles []int, H int) int {
	sum := 0
	for _, p := range piles {
		sum += p
	}
	l := (sum + H - 1) / H
	r := l * 2
	for l < r {
		m := (l + r) / 2
		h := H
		for _, p := range piles {
			h -= (p + m - 1) / m
			if h < 0 {
				break
			}
		}
		if h < 0 {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
