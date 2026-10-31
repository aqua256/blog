package p0032

func longestValidParentheses(s string) int {
	dp := make([]int, len(s)) // dp[i]: 以 s[i] 结尾的最长有效子串长度
	best := 0
	for i := 1; i < len(s); i++ {
		if s[i] != ')' {
			continue // 以 '(' 结尾不可能有效
		}
		if s[i-1] == '(' { // ...()
			dp[i] = 2
			if i >= 2 {
				dp[i] += dp[i-2]
			}
		} else if j := i - dp[i-1] - 1; j >= 0 && s[j] == '(' { // ...((...))
			dp[i] = dp[i-1] + 2
			if j >= 1 {
				dp[i] += dp[j-1] // 接上 j 前面紧挨着的有效串
			}
		}
		best = max(best, dp[i])
	}
	return best
}
