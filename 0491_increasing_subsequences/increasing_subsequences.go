package leetcode0491

// TLE
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	m := make(map[int]int)

	cur := make([][]int, (1<<uint(n))+1)
	cur[0] = []int{}
	curSize := 1

	for _, num := range nums {
		k := curSize
		index := m[num]
		m[num] = k
		for j := index; j < k; j++ {
			if len(cur[j]) > 0 && num < cur[j][len(cur[j])-1] {
				continue
			}

			tmp := append(cur[j], num)
			cur[curSize] = tmp
			curSize++

			if len(tmp) > 1 {
				res = append(res, tmp)
			}
		}
	}
	return res
}
