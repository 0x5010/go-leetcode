package leetcode0407

import "container/heap"

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	if m < 3 {
		return 0
	}
	n := len(heightMap[0])
	if n < 3 {
		return 0
	}

	ch := make(cellHeap, 0, m*2+n*2)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		visited[i][0] = true
		visited[i][n-1] = true
		ch = append(ch, cell{r: i, c: 0, h: heightMap[i][0]})
		ch = append(ch, cell{r: i, c: n - 1, h: heightMap[i][n-1]})
	}
	for j := 0; j < n; j++ {
		visited[0][j] = true
		visited[m-1][j] = true
		ch = append(ch, cell{r: 0, c: j, h: heightMap[0][j]})
		ch = append(ch, cell{r: m - 1, c: j, h: heightMap[m-1][j]})
	}

	heap.Init(&ch)
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	res := 0
	for len(ch) > 0 {
		c := heap.Pop(&ch).(cell)
		for _, d := range dirs {
			i, j := c.r+d[0], c.c+d[1]

			if 0 <= i && i < m && 0 <= j && j < n && !visited[i][j] {
				visited[i][j] = true
				res += max(0, c.h-heightMap[i][j])
				heap.Push(&ch, cell{r: i, c: j, h: max(heightMap[i][j], c.h)})
			}
		}
	}
	return res
}

type cell struct {
	r, c, h int
}

type cellHeap []cell

func (ch cellHeap) Len() int           { return len(ch) }
func (ch cellHeap) Less(i, j int) bool { return ch[i].h < ch[j].h }
func (ch cellHeap) Swap(i, j int)      { ch[i], ch[j] = ch[j], ch[i] }

func (ch *cellHeap) Push(x interface{}) {
	*ch = append(*ch, x.(cell))
}

func (ch *cellHeap) Pop() interface{} {
	old := *ch
	n := len(old)
	x := old[n-1]
	*ch = old[0 : n-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
