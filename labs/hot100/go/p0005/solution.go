package p0005

// 中心扩展：枚举 2n-1 个中心向两边扩，O(n^2) 时间 O(1) 空间
func longestPalindrome(s string) string {
	expand := func(l, r int) (int, int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l, r = l-1, r+1
		}
		return l + 1, r - 1 // 回文区间 [l+1, r-1]
	}
	start, end := 0, 0
	for i := range len(s) {
		l1, r1 := expand(i, i)   // 奇数长度
		l2, r2 := expand(i, i+1) // 偶数长度
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

// 区间 DP：dp[i][j] = s[i]==s[j] && (j-i<2 || dp[i+1][j-1])，i 倒序保证 dp[i+1] 已算
func longestPalindromeDP(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	start, end := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 2 || dp[i+1][j-1])
			if dp[i][j] && j-i > end-start {
				start, end = i, j
			}
		}
	}
	return s[start : end+1]
}
