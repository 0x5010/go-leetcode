package leetcode0423

func originalDigits(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	m := [10]int{}
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'z':
			m[0]++ // 0
		case 'w':
			m[2]++ // 2
		case 'x':
			m[6]++ // 6
		case 'u':
			m[4]++ // 4
		case 'g':
			m[8]++ // 8
		case 's':
			m[7]++ // 7 & 6
		case 'v':
			m[5]++ // 5 & 7
		case 'h':
			m[3]++ // 3 & 8
		case 'i':
			m[9]++ // 9 & 5 & 6 & 8
		case 'o':
			m[1]++ // 1 & 2 & 4 & 0
		}
	}

	m[3] -= m[8]
	m[7] -= m[6]
	m[5] -= m[7]
	m[1] -= m[2] + m[0] + m[4]
	m[9] -= m[5] + m[6] + m[8]

	bs := []byte{}
	for i := 0; i < 10; i++ {
		for m[i] > 0 {
			bs = append(bs, byte(i)+'0')
			m[i]--
		}
	}
	return string(bs)
}
