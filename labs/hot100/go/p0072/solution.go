package p0072

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// dp[i][j]: word1[:i] 变成 word2[:j] 的最少操作数
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i // 删光
	}
	for j := range n + 1 {
		dp[0][j] = j // 全插入
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 替换 / 删除 word1[i-1] / 插入 word2[j-1]
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
