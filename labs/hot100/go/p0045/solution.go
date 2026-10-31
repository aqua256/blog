package p0045

func jump(nums []int) int {
	steps, end, farthest := 0, 0, 0 // end: 当前步数能覆盖的右边界
	for i := 0; i < len(nums)-1; i++ { // 不看最后一格，否则可能多跳一步
		farthest = max(farthest, i+nums[i])
		if i == end { // 走到本层边界，必须再跳一步
			steps++
			end = farthest
		}
	}
	return steps
}
