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
	sortByMerge(tmp)
	res := strings.Join(tmp, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}

func sortByMerge(l []string) {
	sort.Slice(l, func(i, j int) bool {
		return l[i]+l[j] < l[j]+l[i]
	})
}
