package leetcode0806

func numberOfLines(widths []int, S string) []int {
	n := len(S)
	if n == 0 {
		return []int{0, 0}
	}
	lines, sum := 1, 0
	for i := 0; i < n; i++ {
		l := widths[S[i]-'a']
		if sum+l > 100 {
			lines++
			sum = l
		} else {
			sum += l
		}
	}
	return []int{lines, sum}
}
