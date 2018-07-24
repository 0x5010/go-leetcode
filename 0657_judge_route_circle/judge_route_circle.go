package leetcode0657

func judgeCircle(moves string) bool {
	h, v := 0, 0
	for _, move := range moves {
		switch move {
		case 'U':
			h++
		case 'D':
			h--
		case 'L':
			v++
		case 'R':
			v--
		}
	}
	return h == 0 && v == 0
}
