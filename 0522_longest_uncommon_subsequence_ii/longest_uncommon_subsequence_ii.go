package leetcode0522

func findLUSlength(strs []string) int {
	n := len(strs)
	res := -1
	for i, str := range strs {
		if len(str) < res {
			continue
		}
		j := 0
		for j < n {
			if i != j && isSubseq(str, strs[j]) {
				break
			}
			j++
		}
		if j == n {
			res = max(res, len(str))
		}
	}
	return res
}

func isSubseq(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	return i == len(a)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
