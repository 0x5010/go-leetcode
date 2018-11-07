package leetcode0934

func shortestBridge(A [][]int) int {
	n := len(A)

	var paint func(int, int) int
	paint = func(i, j int) int {
		if i < 0 || j < 0 || i == n || j == n || A[i][j] != 1 {
			return 0
		}
		A[i][j] = 2
		return 1 + paint(i-1, j) + paint(i+1, j) + paint(i, j-1) + paint(i, j+1)
	}

	var expand func(int, int, int) bool
	expand = func(i, j, color int) bool {
		if i < 0 || j < 0 || i == n || j == n {
			return false
		}
		if A[i][j] == 0 {
			A[i][j] = color + 1
		}
		return A[i][j] == 1
	}

	for i, found := 0, 0; found == 0 && i < n; i++ {
		for j := 0; found == 0 && j < n; j++ {
			found = paint(i, j)
		}
	}
	for color := 2; ; color++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if A[i][j] == color &&
					((expand(i-1, j, color) || expand(i, j-1, color)) ||
						(expand(i+1, j, color) || expand(i, j+1, color))) {
					return color - 2
				}
			}
		}
	}
}
