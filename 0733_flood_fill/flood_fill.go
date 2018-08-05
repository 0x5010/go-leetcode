package leetcode0733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	m, n := len(image), len(image[0])
	s := [][2]int{{sr, sc}}
	for len(s) != 0 {
		p := s[0]
		s = s[1:]
		image[p[0]][p[1]] = newColor

		for _, dir := range dirs {
			x, y := p[0]+dir[0], p[1]+dir[1]
			if 0 <= x && x < m && 0 <= y && y < n &&
				image[x][y] == oldColor {
				s = append(s, [2]int{x, y})
			}
		}
	}
	return image
}
