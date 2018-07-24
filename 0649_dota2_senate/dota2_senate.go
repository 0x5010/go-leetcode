package leetcode0649

func predictPartyVictory(senate string) string {
	n := len(senate)
	qr := make([]int, 0, n)
	qd := make([]int, 0, n)
	for i, b := range senate {
		if b == 'R' {
			qr = append(qr, i)
		} else {
			qd = append(qd, i)
		}
	}

	for len(qr) > 0 && len(qd) > 0 {
		ri, di := qr[0], qd[0]
		qr, qd = qr[1:], qd[1:]
		if ri < di {
			qr = append(qr, ri+n)
		} else {
			qd = append(qd, di+n)
		}
	}

	if len(qr) > 0 {
		return "Radiant"
	}
	return "Dire"
}
