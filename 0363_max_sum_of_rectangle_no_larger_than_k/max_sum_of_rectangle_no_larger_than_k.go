package leetcode0363

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	res := math.MinInt32
	M, N := min(m, n), max(m, n)

	for p := 0; p < M; p++ {
		tmp := make([]int, N)
		for i := p; i < M; i++ {
			sums := []int{0}
			curSum, curMax := 0, math.MinInt32

			for j := 0; j < N; j++ {
				if m < n {
					tmp[j] += matrix[i][j]
				} else {
					tmp[j] += matrix[j][i]
				}

				sums = append(sums, sums[len(sums)-1]+tmp[j])
				curMax = max(curMax, sums[len(sums)-1]-curSum)
				curSum = min(curSum, sums[len(sums)-1])
			}

			if curMax == k {
				return k
			} else if curMax < k {
				res = max(res, curMax)
			} else {
				tmpRes := search(sums, 0, N+1, k)
				if tmpRes == k {
					return k
				}
				res = max(res, tmpRes)
			}
		}
	}
	return res
}

func search(sums []int, i, j, target int) int {
	if i+1 >= j {
		return math.MinInt32
	}
	m := (i + j) / 2
	res := max(search(sums, i, m, target), search(sums, m, j, target))
	if res == target {
		return res
	}

	l, r := i, m
	for l < m && r < j {
		tmp := sums[r] - sums[l]
		if tmp == target {
			return target
		} else if tmp > target {
			l++
		} else {
			res = max(res, tmp)
			r++
		}
	}
	copy(sums[i:j], merge(sums[i:m], sums[m:j]))
	return res
}

func merge(a, b []int) []int {
	m, n := len(a), len(b)
	tmp := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if a[i] < b[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = b[j]
			j++
		}
		k++
	}
	if i == m {
		copy(tmp[k:], b[j:])
	} else {
		copy(tmp[k:], a[i:])
	}
	return tmp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
