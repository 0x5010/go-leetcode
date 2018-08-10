package leetcode0769

func maxChunksToSorted(arr []int) int {
	max := 0
	res := 0
	for i, v := range arr {
		if v > max {
			max = v
		}
		if max == i {
			res++
		}
	}
	return res
}
