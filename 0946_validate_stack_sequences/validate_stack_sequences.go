package leetcode0946

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	s := make([]int, 0, n)
	i := 0
	for _, x := range pushed {
		s = append(s, x)
		for len(s) != 0 && s[len(s)-1] == popped[i] {
			s = s[:len(s)-1]
			i++
		}
	}
	return i == n
}
