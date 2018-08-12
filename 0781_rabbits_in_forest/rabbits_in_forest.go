package leetcode0781

func numRabbits(answers []int) int {
	count := [1000]int{}
	for _, v := range answers {
		count[v]++
	}
	res := 0
	for v, c := range count {
		if c == 0 {
			continue
		}
		v++
		res += c / v * v
		if c%v > 0 {
			res += v
		}
	}
	return res
}
