package tpl_dp

// DP 四步：
//  1. 状态定义：dp[i] 表示什么（前 i 个 / 以 i 结尾 / 容量为 i）
//  2. 转移：dp[i] 由哪些更小的状态推来（最后一步做了什么）
//  3. 初始值：最小子问题，以及「不可达」用什么表示
//  4. 遍历顺序：保证算 dp[i] 时依赖项已算好；只依赖前几项就滚动

// 一维线性：dp[i] = max(dp[i-1], dp[i-2]+nums[i])（打家劫舍型），滚动成两个变量
func linear(nums []int) int {
	prev, cur := 0, 0 // dp[i-2], dp[i-1]
	for _, x := range nums {
		prev, cur = cur, max(cur, prev+x)
	}
	return cur
}

// 0-1 背包：每件物品最多选一次。dp[j] = 容量 j 下的最大价值
// 容量倒序遍历，保证 dp[j-w] 还是「上一件物品」时的值
func knapsack01(weights, values []int, capacity int) int {
	dp := make([]int, capacity+1)
	for i, w := range weights {
		for j := capacity; j >= w; j-- {
			dp[j] = max(dp[j], dp[j-w]+values[i])
		}
	}
	return dp[capacity]
}

// 完全背包（最少个数）：每件物品无限次。dp[j] = 凑出 j 的最少件数，inf 表示凑不出
// 容量正序遍历，dp[j-w] 可能已经用过这件物品，正是「无限次」
func knapsackCompleteMin(weights []int, target int) int {
	inf := target + 1
	dp := make([]int, target+1)
	for j := 1; j <= target; j++ {
		dp[j] = inf
	}
	for _, w := range weights {
		for j := w; j <= target; j++ {
			dp[j] = min(dp[j], dp[j-w]+1)
		}
	}
	if dp[target] == inf {
		return -1
	}
	return dp[target]
}
