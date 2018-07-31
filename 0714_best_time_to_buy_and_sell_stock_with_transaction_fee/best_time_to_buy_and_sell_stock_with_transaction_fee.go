package leetcode0714

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	res := 0
	hold := 0 - prices[0] - fee
	for i := 1; i < n; i++ {
		tmp := res
		res = max(res, hold+prices[i])
		hold = max(hold, tmp-prices[i]-fee)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
