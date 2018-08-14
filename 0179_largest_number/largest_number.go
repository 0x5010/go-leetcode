package leetcode0179

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	tmp := make([]string, len(nums))
	for i := range tmp {
		tmp[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i]+tmp[j] > tmp[j]+tmp[i]
	})
	res := strings.Join(tmp, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}
