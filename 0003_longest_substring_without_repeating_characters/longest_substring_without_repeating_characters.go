package leetcode0003

func lengthOfLongestSubstring(s string) int {
	index := [256]int{} // len for Extended ASCII
	for i := range index {
		index[i] = -1
	}

	maxLen, i := 0, -1

	for j, c := range s {
		if index[c] > i {
			i = index[c]
		}
		l := j - i
		if l > maxLen {
			maxLen = l
		}
		index[c] = j
	}
	return maxLen
}
