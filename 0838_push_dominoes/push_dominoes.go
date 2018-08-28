package leetcode0838

func pushDominoes(dominoes string) string {
	d := "L" + dominoes + "R"
	n := len(d)
	bs := make([]byte, 0, n-2)
	for i, j := 0, 1; j < n; j++ {
		if d[j] == '.' {
			continue
		}
		if i != 0 {
			bs = append(bs, d[i])
		}
		m := j - i - 1
		if d[i] == d[j] {
			for k := 0; k < m; k++ {
				bs = append(bs, d[i])
			}
		} else if d[i] == 'L' && d[j] == 'R' {
			for k := 0; k < m; k++ {
				bs = append(bs, '.')
			}
		} else {
			for k := 0; k < m/2; k++ {
				bs = append(bs, 'R')
			}
			if m%2 == 1 {
				bs = append(bs, '.')
			}
			for k := 0; k < m/2; k++ {
				bs = append(bs, 'L')
			}
		}
		i = j
	}
	return string(bs)
}
