package leetcode0936

func movesToStamp(stamp string, target string) []int {
	m, n := len(stamp), len(target)
	count := 0
	bStamp, bTarget := []byte(stamp), []byte(target)
	unvisited := map[int]struct{}{}
	for i := 0; i <= n-m; i++ {
		unvisited[i] = struct{}{}
	}
	res := []int{}
L:
	for count < n {
		for index := range unvisited {
			found := false
			var i, j int
			for i, j = index, 0; j < m; i, j = i+1, j+1 {
				if bTarget[i] != '*' && bTarget[i] != bStamp[j] {
					break
				} else if bTarget[i] == bStamp[j] {
					found = true
				}
			}
			if j == m && found {
				delete(unvisited, index)
				for k := index; k < index+m; k++ {
					if bTarget[k] != '*' {
						bTarget[k] = '*'
						count++
					}
				}
				res = append(res, index)
				continue L
			}
		}
		return []int{}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
