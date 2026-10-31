package p0300

import "slices"

// 标准 DP：dp[i] = 以 nums[i] 结尾的最长递增子序列长度，O(n^2)
func lengthOfLISDP(nums []int) int {
	dp := make([]int, len(nums))
	best := 0
	for i := range nums {
		dp[i] = 1
		for j := range i {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		best = max(best, dp[i])
	}
	return best
}

// 贪心 + 二分：tails[k] = 长度为 k+1 的递增子序列的最小结尾，O(n log n)
func lengthOfLIS(nums []int) int {
	var tails []int
	for _, x := range nums {
		i, _ := slices.BinarySearch(tails, x) // 第一个 >= x 的位置
		if i == len(tails) {
			tails = append(tails, x)
		} else {
			tails[i] = x
		}
	}
	return len(tails)
}
