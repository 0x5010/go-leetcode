package leetcode0228

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return nil
	}

	res := []string{}
	begin := nums[0]
	s := ""
	for i := 0; i < n; i++ {
		if i == n-1 || nums[i]+1 != nums[i+1] {
			if nums[i] == begin {
				s = fmt.Sprintf("%d", begin)
			} else {
				s = fmt.Sprintf("%d->%d", begin, nums[i])
			}

			res = append(res, s)
			if i+1 < n {
				begin = nums[i+1]
			}
		}
	}
	return res
}
