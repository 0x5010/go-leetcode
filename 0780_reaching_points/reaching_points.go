package leetcode0780

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx >= sx && ty >= sy {
		if tx > ty {
			if sy == ty {
				return (tx-sx)%ty == 0
			}
			tx %= ty
		} else {
			if sx == tx {
				return (ty-sy)%tx == 0
			}
			ty %= tx
		}
	}
	return false
}
