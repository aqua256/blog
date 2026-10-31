package p0121

func maxProfit(prices []int) int {
	minPrice, best := prices[0], 0
	for _, p := range prices {
		minPrice = min(minPrice, p) // 截至今天的最低买入价
		best = max(best, p-minPrice)
	}
	return best
}
