package leetcode0967

func numsSameConsecDiff(N int, K int) []int {
	cur := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	tmp := make([]int, 0, 10)
	for i := 2; i <= N; i++ {
		for _, x := range cur {
			y := x % 10
			if x > 0 && y+K < 10 {
				tmp = append(tmp, x*10+y+K)
			}
			if x > 0 && K > 0 && y-K >= 0 {
				tmp = append(tmp, x*10+y-K)
			}
		}
		cur, tmp = tmp, cur
		tmp = tmp[:0]
	}
	return cur
}
