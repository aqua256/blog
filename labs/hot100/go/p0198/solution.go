package p0198

func rob(nums []int) int {
	// dp[i] = max(dp[i-1], dp[i-2]+nums[i])：第 i 家偷或不偷
	prev, cur := 0, 0 // dp[i-2], dp[i-1]
	for _, x := range nums {
		prev, cur = cur, max(cur, prev+x)
	}
	return cur
}
