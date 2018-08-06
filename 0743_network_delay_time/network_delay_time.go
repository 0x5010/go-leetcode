package leetcode0743

func networkDelayTime(times [][]int, N int, K int) int {
	n := len(times)
	edges := make([][][]int, N+1)
	d := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		d[i] = -1
		edges[i] = [][]int{}
	}
	d[K] = 0
	for i := 0; i < n; i++ {
		edges[times[i][0]] = append(edges[times[i][0]], []int{times[i][1], times[i][2]})
	}
	done := map[int]bool{}
	res := 0
	for i := 1; i < N+1; i++ {
		choose := -1
		for j := 1; j < N+1; j++ {
			if d[j] != -1 && !done[j] && (choose == -1 || d[j] < d[choose]) {
				choose = j
			}
		}
		if choose == -1 {
			return -1
		}
		done[choose] = true
		for j := 0; j < len(edges[choose]); j++ {
			index := edges[choose][j][0]
			tmp := edges[choose][j][1] + d[choose]
			if d[index] == -1 || d[index] > tmp {
				d[index] = tmp
			}
		}
		if d[choose] > res {
			res = d[choose]
		}
	}
	return res
}
