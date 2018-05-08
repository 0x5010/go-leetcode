package leetcode0389

func findTheDifference(s string, t string) byte {
	var res byte
	for i := 0; i < len(s); i++ {
		res ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		res ^= t[i]
	}
	return res
}
