package leetcode0587

/**
 * Definition for a point.
 * type Point struct {
 *     X int
 *     Y int
 * }
 */

func outerTrees(points []Point) []Point {
	n := len(points)
	if n < 4 {
		return points
	}

	first := 0
	for i := 1; i < n; i++ {
		if points[i].X < points[first].X ||
			(points[i].X == points[first].X && points[i].Y < points[first].Y) {
			first = i
		}
	}

	res := []Point{points[first]}
	cur := first
	for {
		next := (cur + 1) % n
		for i := 0; i < n; i++ {
			if i == cur {
				continue
			}
			cross := crossProductLength(points[cur], points[next], points[i])
			if cross < 0 ||
				(cross == 0 && isFurther(points[cur], points[next], points[i])) {
				next = i
			}
		}

		for i := 0; i < n; i++ {
			if i == cur || i == first {
				continue
			}
			if crossProductLength(points[cur], points[next], points[i]) == 0 {
				res = append(res, points[i])
			}
		}

		cur = next

		if cur == first {
			break
		}
	}

	return res
}

func crossProductLength(o, a, b Point) int {
	return (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
}

func isFurther(o, a, b Point) bool {
	return distance(o, a) < distance(o, b)
}

func distance(a, b Point) int {
	return (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)
}
