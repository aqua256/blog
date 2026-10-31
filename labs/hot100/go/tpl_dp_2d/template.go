package tpl_dp_2d

// 网格 DP：dp[i][j] 只依赖上方和左方，按行从上到下、从左到右；可压成一行
func gridMinPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = grid[0][0]
			case i == 0:
				dp[i][j] = dp[i][j-1] + grid[i][j]
			case j == 0:
				dp[i][j] = dp[i-1][j] + grid[i][j]
			default:
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

// 双串 DP：dp[i][j] 表示 a[:i] 与 b[:j] 的答案，多开一行一列放空串
// 看 a[i-1] 和 b[j-1]：相等走对角线，否则从左/上/左上里取最优
func twoStrings(a, b string) int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1) // 初始值：空串对任何串 LCS 为 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// 区间 DP：dp[i][j] 依赖更短的区间 dp[i+1][j-1]，所以 i 倒序、j 正序
func intervalPalindromes(s string) int {
	n, count := len(s), 0
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 2 || dp[i+1][j-1])
			if dp[i][j] {
				count++ // 回文子串个数
			}
		}
	}
	return count
}
