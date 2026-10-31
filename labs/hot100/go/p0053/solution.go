package p0053

func maxSubArray(nums []int) int {
	cur, best := nums[0], nums[0] // cur：以当前元素结尾的最大子数组和
	for _, x := range nums[1:] {
		cur = max(cur+x, x) // 前面的和为负就丢掉，从自己重新开始
		best = max(best, cur)
	}
	return best
}
