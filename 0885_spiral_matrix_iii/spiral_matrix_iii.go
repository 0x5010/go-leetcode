package leetcode0885

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	move := R*C - 1
	x, y, dx, dy, round := r0, c0, 0, 1, 2
	res := [][]int{[]int{x, y}}
	for move > 0 {
		for m := round / 2; m > 0; m-- {
			x += dx
			y += dy
			if 0 <= x && x < R && 0 <= y && y < C {
				res = append(res, []int{x, y})
				move--
			}
		}
		round++
		dx, dy = dy, -dx
	}
	return res
}
