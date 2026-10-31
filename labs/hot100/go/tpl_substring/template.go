package tpl_substring

// 前缀和 + 哈希：子数组 [j, i) 的和 = pre[i] - pre[j]。
// 枚举右端点 i，在哈希表里查「之前出现过多少个 pre[j] = pre[i] - k」。
// 不要求非负，滑动窗口失效（有负数）时用它。
func countSubarrays(nums []int, k int) int {
	cnt := map[int]int{0: 1} // 必须预置空前缀
	sum, ans := 0, 0
	for _, x := range nums {
		sum += x
		ans += cnt[sum-k] // 先查后存：避免空子数组
		cnt[sum]++
	}
	return ans
}

// 单调队列：维护窗口最值。队列存下标，值单调；入队时弹掉被支配的，队头过期就出队。
func windowMax(nums []int, k int) []int {
	var dq, ans []int
	for i, x := range nums {
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= x {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		if dq[0] <= i-k {
			dq = dq[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[dq[0]])
		}
	}
	return ans
}
