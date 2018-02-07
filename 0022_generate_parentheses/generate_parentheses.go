package leetcode0022

func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		res = append(res, "")
	} else {
		for i := 0; i < n; i++ {
			for _, l := range generateParenthesis(i) {
				for _, r := range generateParenthesis(n - i - 1) {
					res = append(res, "("+l+")"+r)
				}
			}
		}
	}
	return res
}
