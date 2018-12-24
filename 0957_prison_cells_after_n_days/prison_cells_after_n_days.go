package leetcode0957

func prisonAfterNDays(cells []int, N int) []int {
	for N = (N-1)%14 + 1; N > 0; N-- {
		tmp := make([]int, 8)
		for i := 1; i < 7; i++ {
			if cells[i-1] == cells[i+1] {
				tmp[i] = 1
			}
		}
		cells = tmp
	}
	return cells
}
