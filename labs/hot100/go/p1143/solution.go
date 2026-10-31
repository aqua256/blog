package p1143

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// dp[i][j]: text1[:i] 与 text2[:j] 的 LCS 长度；多一行一列免去边界判断
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
