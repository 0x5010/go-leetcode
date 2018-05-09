package leetcode0393

func validUtf8(data []int) bool {
	count := 0
	for _, i := range data {
		if count == 0 {
			if i>>3 == 30 { // 0b11110
				count = 3
			} else if i>>4 == 14 { // 0b1110
				count = 2
			} else if i>>5 == 6 { // 0b110
				count = 1
			} else if i>>7 > 0 {
				return false
			}
		} else {
			if i>>6 != 2 {
				return false
			}
			count--
		}
	}
	return count == 0
}
