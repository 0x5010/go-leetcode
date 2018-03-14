package leetcode0164

import (
	"math"
)

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	max, min := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	gap := int(math.Ceil(float64(max-min) / float64(n-1)))
	bucketsMin := make([]int, n-1)
	for i := range bucketsMin {
		bucketsMin[i] = math.MaxInt32
	}
	bucketsMax := make([]int, n-1)
	for i := range bucketsMax {
		bucketsMax[i] = math.MinInt32
	}

	for _, num := range nums {
		if num == max || num == min {
			continue
		}
		index := (num - min) / gap
		if num < bucketsMin[index] {
			bucketsMin[index] = num
		}
		if num > bucketsMax[index] {
			bucketsMax[index] = num
		}
	}
	pre := min
	res := math.MinInt32
	for i := 0; i < n-1; i++ {
		if bucketsMin[i] == math.MaxInt32 && bucketsMax[i] == math.MinInt32 {
			continue
		}
		if bucketsMin[i]-pre > res {
			res = bucketsMin[i] - pre
		}
		pre = bucketsMax[i]
	}
	if max-pre > res {
		res = max - pre
	}
	return res
}
