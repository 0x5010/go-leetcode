package leetcode0735

func asteroidCollision(asteroids []int) []int {
	s := []int{}
	for _, asteroid := range asteroids {
		for len(s) != 0 && s[len(s)-1] > 0 && s[len(s)-1] < -asteroid {
			s = s[:len(s)-1]
		}
		if len(s) == 0 || asteroid > 0 || s[len(s)-1] < 0 {
			s = append(s, asteroid)
		} else if asteroid < 0 && s[len(s)-1] == -asteroid {
			s = s[:len(s)-1]
		}
	}
	return s
}
