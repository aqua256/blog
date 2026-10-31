package p0239

func maxSlidingWindow(nums []int, k int) []int {
	var dq []int // 存下标，对应值单调递减；队头即窗口最大值
	ans := make([]int, 0, len(nums)-k+1)
	for i, x := range nums {
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= x {
			dq = dq[:len(dq)-1] // 比新来的小的永远不会再当最大值
		}
		dq = append(dq, i)
		if dq[0] <= i-k {
			dq = dq[1:] // 队头滑出窗口
		}
		if i >= k-1 {
			ans = append(ans, nums[dq[0]])
		}
	}
	return ans
}
