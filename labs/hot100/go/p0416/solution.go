package p0416

func canPartition(nums []int) bool {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum%2 == 1 {
		return false
	}
	// 0-1 背包：能否选出若干数，和恰为 sum/2
	target := sum / 2
	dp := make([]bool, target+1) // dp[j]: 能否凑出和 j
	dp[0] = true
	for _, x := range nums {
		for j := target; j >= x; j-- { // 倒序，保证每个数只用一次
			dp[j] = dp[j] || dp[j-x]
		}
	}
	return dp[target]
}
