package leetcode0216

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	var dfs func(int, int, int, []int)
	dfs = func(level, start, n int, path []int) {
		if level == k && n == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n && i < 10; i++ {
			path = append(path, i)
			dfs(level+1, i+1, n-i, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 1, n, []int{})
	return res
}
