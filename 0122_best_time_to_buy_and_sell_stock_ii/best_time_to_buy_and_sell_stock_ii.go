package leetcode0122

func maxProfit(prices []int) int {
	res := 0
	for i, p := range prices {
		if i == 0 {
			continue
		}
		if p > prices[i-1] {
			res += p - prices[i-1]
		}
	}
	return res
}
