package leetcode0886

func possibleBipartition(N int, dislikes [][]int) bool {
	dislike := make([][]int, N+1)
	for _, d := range dislikes {
		dislike[d[0]] = append(dislike[d[0]], d[1])
		dislike[d[1]] = append(dislike[d[1]], d[0])
	}
	colors := make([]int, N+1)
	color := 1
	queue := []int{}
	for i := 1; i < N+1; i++ {
		if colors[i] != 0 {
			continue
		}
		queue = append(queue, i)
		colors[i] = color
		for len(queue) != 0 {
			color = -color
			n := len(queue)
			for j := 0; j < n; j++ {
				q := queue[j]
				for _, p := range dislike[q] {
					if colors[p] == 0 {
						colors[p] = color
						queue = append(queue, p)
					} else if colors[p] != color {
						return false
					}
				}
			}
			queue = queue[n:]
		}
	}
	return true
}
