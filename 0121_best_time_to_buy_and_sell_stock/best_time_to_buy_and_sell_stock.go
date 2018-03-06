package leetcode0121

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max, minElement := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minElement {
			minElement = prices[i]
		}
		tmp := prices[i] - minElement
		if tmp > max {
			max = tmp
		}
	}
	return max
}
