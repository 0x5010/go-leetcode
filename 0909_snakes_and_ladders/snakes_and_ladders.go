package leetcode0909

func snakesAndLadders(board [][]int) int {
	n := len(board)
	nn := n * n
	dist := map[int]int{1: 0}
	queue := []int{1}
	for len(queue) != 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			cur := queue[i]
			if cur >= nn {
				return dist[cur]
			}
			for j := 1; j < 7; j++ {
				index := cur + j
				if index > nn {
					break
				}
				x, y := getCoord(n, index)
				next := board[x][y]
				if next == -1 {
					next = index
				}
				if _, ok := dist[next]; !ok {
					dist[next] = dist[cur] + 1
					queue = append(queue, next)
				}
			}
		}
		queue = queue[count:]
	}
	return -1
}

func getCoord(n, i int) (int, int) {
	i--
	x, y := n-i/n-1, i%n
	if (n-x)%2 == 0 {
		y = n - 1 - y
	}
	return x, y
}
