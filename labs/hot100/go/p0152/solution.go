package p0152

func maxProduct(nums []int) int {
	// 同时维护以当前元素结尾的最大积和最小积：负数会让两者互换
	best, curMax, curMin := nums[0], nums[0], nums[0]
	for _, x := range nums[1:] {
		if x < 0 {
			curMax, curMin = curMin, curMax
		}
		curMax = max(x, curMax*x)
		curMin = min(x, curMin*x)
		best = max(best, curMax)
	}
	return best
}
