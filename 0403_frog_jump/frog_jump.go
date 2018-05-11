package leetcode0403

func canCross(stones []int) bool {
	n := len(stones)
	if n == 0 || stones[1] != 1 {
		return false
	}

	if n == 1 || n == 2 {
		return true
	}

	last := stones[n-1]
	m := make(map[int]struct{}, n)

	for i := 0; i < n; i++ {
		if i > 3 && stones[i] > stones[i-1]*2 {
			return false
		}
		m[stones[i]] = struct{}{}
	}

	var dfs func(int, int) bool
	dfs = func(index, jump int) bool {
		tmp := index + jump
		if tmp-1 == last || tmp == last || tmp+1 == last {
			return true
		}
		for i := 1; -1 <= i; i-- {
			if jump+i > 0 {
				if _, ok := m[tmp+i]; ok && dfs(tmp+i, jump+i) {
					return true
				}
			}
		}
		return false
	}
	return dfs(1, 1)
}
