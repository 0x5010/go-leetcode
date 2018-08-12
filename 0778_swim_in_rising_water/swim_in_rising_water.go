package leetcode0778

import "container/heap"

func swimInWater(grid [][]int) int {
	n := len(grid)
	pq := PQ{}
	visited := make([]bool, n*n)
	heap.Push(&pq, &entry{
		x:   0,
		y:   0,
		max: grid[0][0],
	})
	visited[0] = true
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for {
		pre := heap.Pop(&pq).(*entry)
		for _, dir := range dirs {
			x, y := pre.x+dir[0], pre.y+dir[1]
			if x < 0 || x >= n || y < 0 || y >= n || visited[x*n+y] {
				continue
			}
			if x == n-1 && y == n-1 {
				return max(pre.max, grid[x][y])
			}
			heap.Push(&pq, &entry{
				x:   x,
				y:   y,
				max: max(pre.max, grid[x][y]),
			})
			visited[x*n+y] = true
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type entry struct {
	x, y int
	max  int
}

type PQ []*entry

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].max < pq[j].max }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQ) Push(x interface{}) {
	tmp := x.(*entry)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	tmp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return tmp
}
