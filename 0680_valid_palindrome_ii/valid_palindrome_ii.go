package leetcode0680

func validPalindrome(s string) bool {
	bs := []byte(s)
	var recur func(int, int) bool
	recur = func(l, r int) bool {
		for l < r {
			if bs[l] != bs[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if bs[l] == bs[r] {
			continue
		}
		return recur(l+1, r) || recur(l, r-1)
	}
	return true
}
