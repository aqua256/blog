package p0041

func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 原地哈希：把值 v (1<=v<=n) 换到下标 v-1；答案必在 [1, n+1]
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= n && i != nums[i]-1 {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return n + 1
}
