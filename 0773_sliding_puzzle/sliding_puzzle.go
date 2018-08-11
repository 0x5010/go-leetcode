package leetcode0773

func slidingPuzzle(board [][]int) int {
	origin, target := [6]int{1, 2, 3, 4, 5, 0}, [6]int{}
	for i := 0; i < 6; i++ {
		target[i] = board[i/3][i%3]
	}

	if origin == target {
		return 0
	}

	queue := [][6]int{origin}

	seen := map[[6]int]struct{}{}
	seen[origin] = struct{}{}

	res := 1
	dirs := []int{-1, 1, 3, -3}
	count := 1
	for len(queue) > 0 {
		s := queue[0]
		i := -1
		for index := 0; index < 6; index++ {
			if s[index] == 0 {
				i = index
			}
		}
		for _, dir := range dirs {

			j := i + dir
			if j < 0 ||
				j > 5 ||
				i == 2 && j == 3 ||
				i == 3 && j == 2 {
				continue
			}
			c := s
			c[i], c[j] = c[j], c[i]
			if c == target {
				return res
			}

			if _, ok := seen[c]; !ok {
				queue = append(queue, c)
				seen[c] = struct{}{}
			}
		}

		queue = queue[1:]
		count--
		if count == 0 {
			count = len(queue)
			res++
		}
	}
	return -1
}
