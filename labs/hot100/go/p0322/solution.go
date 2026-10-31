package p0322

func coinChange(coins []int, amount int) int {
	// dp[i]: 凑出 i 的最少硬币数；amount+1 当作无穷大
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if c <= i {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
