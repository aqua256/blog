package p0279

func numSquares(n int) int {
	// 完全背包：物品是 1,4,9,...，每个可无限用，求凑出 n 的最少个数
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i // 全用 1
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
