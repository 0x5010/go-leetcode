package leetcode0405

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	h := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f"}
	res := ""
	for i := 0; i < 8 && num != 0; i++ {
		res = h[num&15] + res
		num >>= 4
	}
	return res
}
