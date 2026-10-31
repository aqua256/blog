package p0064

func minPathSum(grid [][]int) int {
	n := len(grid[0])
	dp := make([]int, n) // 滚动一行：dp[j] 更新前是上方的值
	for i, row := range grid {
		for j, v := range row {
			switch {
			case i == 0 && j == 0:
				dp[j] = v
			case i == 0:
				dp[j] = dp[j-1] + v // 第一行只能从左来
			case j == 0:
				dp[j] += v // 第一列只能从上来
			default:
				dp[j] = min(dp[j], dp[j-1]) + v
			}
		}
	}
	return dp[n-1]
}
