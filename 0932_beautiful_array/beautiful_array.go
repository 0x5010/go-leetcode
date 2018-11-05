package leetcode0932

func beautifulArray(N int) []int {
	res := []int{1}
	tmp := []int{}
	for len(res) < N {
		for _, i := range res {
			if i*2-1 <= N {
				tmp = append(tmp, i*2-1)
			}
		}
		for _, i := range res {
			if i*2 <= N {
				tmp = append(tmp, i*2)
			}
		}
		res, tmp = tmp, res
		tmp = tmp[:0]
	}
	return res
}
