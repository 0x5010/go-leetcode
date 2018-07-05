package leetcode0517

func findMinMoves(machines []int) int {
	n := len(machines)
	sum := 0
	for _, v := range machines {
		sum += v
	}
	if sum%n != 0 {
		return -1
	}
	avg, count := sum/n, 0
	res := 0
	for _, v := range machines {
		diff := v - avg
		count += diff
		res = max(max(res, abs(count)), diff)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
