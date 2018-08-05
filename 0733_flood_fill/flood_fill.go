package leetcode0733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	m, n := len(image), len(image[0])
	var fill func(int, int)
	fill = func(r, c int) {
		image[r][c] = newColor

		for _, dir := range dirs {
			x, y := r+dir[0], c+dir[1]
			if 0 <= x && x < m && 0 <= y && y < n &&
				image[x][y] == oldColor {
				fill(x, y)
			}
		}
	}
	fill(sr, sc)
	return image
}
