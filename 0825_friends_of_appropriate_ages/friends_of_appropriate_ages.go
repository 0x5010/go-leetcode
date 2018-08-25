package leetcode0825

func numFriendRequests(ages []int) int {
	count := [121]int{}
	for _, age := range ages {
		count[age]++
	}
	sum, min := 0, 15
	res := 0
	for i := 15; i <= 120; i++ {
		for min <= i/2+7 {
			sum -= count[min]
			min++
		}
		sum += count[i]
		res += count[i] * (sum - 1)
	}
	return res
}
