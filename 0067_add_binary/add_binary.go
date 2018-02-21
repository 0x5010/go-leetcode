package leetcode0067

import "strconv"

func addBinary(a string, b string) string {
	s := ""
	c, i, j := 0, len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		aNum, bNum := 0, 0
		if i >= 0 {
			aNum = int(a[i]) - int('0')
			i--
		}
		if j >= 0 {
			bNum = int(b[j]) - int('0')
			j--
		}
		cNum := aNum + bNum + c
		s = strconv.Itoa(cNum%2) + s
		c = cNum / 2
	}

	if c == 0 {
		return s
	}
	return "1" + s
}
