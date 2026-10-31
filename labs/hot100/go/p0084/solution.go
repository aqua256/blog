package p0084

// 以每根柱子为最矮者向两侧扩展：宽度由「左侧第一个更矮」和「右侧第一个更矮」决定。
// 单调递增栈中，弹出 top 时：当前 i 是它右侧第一个更矮，新栈顶是它左侧第一个更矮。
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0) // 末尾哨兵 0：保证最后栈里剩下的柱子都被弹出结算
	stack := []int{-1}           // 哨兵 -1 作为左边界
	best := 0
	for i, h := range heights {
		for len(stack) > 1 && heights[stack[len(stack)-1]] >= h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			best = max(best, heights[top]*(i-left-1))
		}
		stack = append(stack, i)
	}
	return best
}
