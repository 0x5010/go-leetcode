package leetcode0684

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	ls := make([]int, n)
	zs := make([]int, n)
	for _, e := range edges {
		i := e[0] - 1
		for ls[i] != 0 {
			i = ls[i] - 1
		}
		j := e[1] - 1
		for ls[j] != 0 {
			j = ls[j] - 1
		}
		if i == j {
			return e
		}
		if zs[i] >= zs[j] {
			zs[i] += zs[j] + 1
			ls[j] = i + 1
		} else {
			zs[j] += zs[i] + 1
			ls[i] = j + 1
		}
	}
	return nil
}
