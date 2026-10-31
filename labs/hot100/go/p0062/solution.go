package p0062

func uniquePaths(m int, n int) int {
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]，压成一行：dp[j] 更新前是上方，dp[j-1] 是左方
	dp := make([]int, n)
	for j := range dp {
		dp[j] = 1 // 第一行只有一种走法
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
