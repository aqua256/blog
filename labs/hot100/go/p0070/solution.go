package p0070

func climbStairs(n int) int {
	// dp[i] = dp[i-1] + dp[i-2]，只依赖前两项，滚动成两个变量
	prev, cur := 1, 1 // dp[0], dp[1]
	for range n - 1 {
		prev, cur = cur, prev+cur
	}
	return cur
}
