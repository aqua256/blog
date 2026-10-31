package p0139

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
	}
	dp := make([]bool, len(s)+1) // dp[i]: 前缀 s[:i] 能被拆分
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ { // 枚举最后一个单词 s[j:i]
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
