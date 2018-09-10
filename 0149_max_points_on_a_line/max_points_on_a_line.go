package leetcode0149

/**
 * Definition for a point.
 * type Point struct {
 *     X int
 *     Y int
 * }
 */

func maxPoints(points []Point) int {
	res := 0
	for i := 0; i < len(points); i++ {
		lines := map[Line]int{}
		overlap := 0
		count := 0

		for j := i + 1; j < len(points); j++ {
			l := reduceSegment(points[i], points[j])
			if l.a == 0 && l.b == 0 {
				overlap++
			} else {
				lines[l]++
				if count < lines[l] {
					count = lines[l]
				}
			}
		}
		count += overlap + 1
		if count > res {
			res = count
		}
	}
	return res
}

type Line struct {
	a int
	b int
}

func reduceSegment(p1, p2 Point) Line {
	a := p2.X - p1.X
	b := p2.Y - p1.Y

	if a != 0 && b != 0 {
		r := gcd(a, b)
		a /= r
		b /= r
	} else if a != 0 {
		a = 1
	} else if b != 0 {
		b = 1
	}
	return Line{a, b}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
