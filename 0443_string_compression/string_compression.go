package leetcode0443

import "strconv"

func compress(chars []byte) int {
	res, i, n := 0, 0, len(chars)
	for i < n {
		b := chars[i]
		count := 0
		for i < n && chars[i] == b {
			i++
			count++
		}
		chars[res] = b
		res++
		if count != 1 {
			cs := strconv.Itoa(count)
			for j := 0; j < len(cs); j++ {
				chars[res] = cs[j]
				res++
			}
		}
	}
	return res
}
