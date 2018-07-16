package leetcode0593

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	d12, d13, d14 := d(p1, p2), d(p1, p3), d(p1, p4)
	d23, d24 := d(p2, p3), d(p2, p4)
	d34 := d(p3, p4)
	if d12 == d13 {
		return d12 != 0 && d14 == d23 && d24 == d34 && d12 == d24
	}
	if d12 == d14 {
		return d12 != 0 && d13 == d24 && d23 == d34 && d12 == d23
	}
	if d13 == d14 {
		return d13 != 0 && d12 == d34 && d23 == d24 && d13 == d23
	}
	return false
}

func d(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}
