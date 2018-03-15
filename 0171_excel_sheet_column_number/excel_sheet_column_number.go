package leetcode0171

func titleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		tmp := int(s[i] - 'A' + 1)
		res = res*26 + tmp
	}
	return res
}
