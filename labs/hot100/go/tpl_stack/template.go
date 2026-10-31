package tpl_stack

// 单调栈：栈里存下标，从栈底到栈顶对应的值单调。
// 新元素到来时，把「被它终结」的栈顶全部弹出——弹出那一刻，
// 当前元素就是栈顶的「右侧第一个更大（或更小）」，新栈顶就是它的「左侧第一个」。
// 每个下标进出栈各一次，整体 O(n)。

// nextGreater 返回每个位置右侧第一个严格更大元素的下标，不存在为 -1。
func nextGreater(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{} // 下标，对应值单调不增
	for i, x := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < x {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i // x 是 top 右侧第一个更大的
		}
		stack = append(stack, i)
	}
	return ans
}

// prevSmaller 返回每个位置左侧第一个严格更小元素的下标，不存在为 -1。
// 求「左侧」时，答案就是入栈前清理完毕后的栈顶。
func prevSmaller(nums []int) []int {
	ans := make([]int, len(nums))
	stack := []int{} // 下标，对应值严格递增
	for i, x := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= x {
			stack = stack[:len(stack)-1]
		}
		ans[i] = -1
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
