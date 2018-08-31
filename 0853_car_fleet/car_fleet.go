package leetcode0853

import "sort"

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([]*car, n)
	for i := 0; i < n; i++ {
		cars[i] = &car{
			pos:  position[i],
			time: float64(target-position[i]) / float64(speed[i]),
		}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	cur := 0.
	res := 0
	for _, car := range cars {
		if car.time > cur {
			cur = car.time
			res++
		}
	}
	return res
}

type car struct {
	pos  int
	time float64
}
