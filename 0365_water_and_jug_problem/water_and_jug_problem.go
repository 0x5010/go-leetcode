package leetcode0365

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}

	if x+y < z {
		return false
	}

	if x > y {
		x, y = y, x
	}

	if x == 0 {
		return y == z
	}

	for y%x != 0 {
		x, y = y%x, x
	}
	return z%x == 0
}
