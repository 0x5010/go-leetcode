package leetcode0913

func catMouseGame(graph [][]int) int {
	n := len(graph)
	colors := make([][][2]int, n)
	outdegree := make([][][2]int, n)
	queue := [][4]int{}
	push := func(cat, mouse, move, color int) {
		colors[cat][mouse][move] = color
		queue = append(queue, [4]int{cat, mouse, move, color})
	}

	for i := 0; i < n; i++ {
		colors[i] = make([][2]int, n)
		outdegree[i] = make([][2]int, n)
		for j := 0; j < n; j++ {
			outdegree[i][j][0] = len(graph[j])
			for _, k := range graph[i] {
				if k != 0 {
					outdegree[i][j][1]++
				}
			}
		}
		if i != 0 {
			for k := 0; k < 2; k++ {
				push(i, 0, k, 1)
				push(i, i, k, 2)
			}
		}
	}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		cat, mouse, move, color := cur[0], cur[1], cur[2], cur[3]
		if cat == 2 && mouse == 1 && move == 0 {
			return color
		}
		prevMove := 1 - move
		var turn int
		if prevMove == 1 {
			turn = cat
		} else {
			turn = mouse
		}
		for _, prev := range graph[turn] {
			var prevCat, prevMouse int
			if prevMove == 1 {
				prevCat = prev
				prevMouse = mouse
			} else {
				prevCat = cat
				prevMouse = prev
			}
			if prevCat == 0 {
				continue
			}
			if colors[prevCat][prevMouse][prevMove] > 0 {
				continue
			}
			outdegree[prevCat][prevMouse][prevMove]--
			if prevMove == 1 && color == 2 ||
				prevMove == 0 && color == 1 ||
				outdegree[prevCat][prevMouse][prevMove] == 0 {
				push(prevCat, prevMouse, prevMove, color)
			}

		}
	}
	return colors[2][1][0]
}
