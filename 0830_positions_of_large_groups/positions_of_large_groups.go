package leetcode0830

func largeGroupPositions(S string) [][]int {
	res := [][]int{}
	for l, r := 0, 1; r < len(S); r++ {
		if S[l] != S[r] {
			l = r
			continue
		}
		if r-l+1 == 3 {
			res = append(res, []int{l, r})
		} else if r-l+1 > 3 {
			res[len(res)-1][1] = r
		}
	}
	return res
}
