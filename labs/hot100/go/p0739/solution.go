package p0739

// 单调栈（栈底到栈顶温度不增）：今天比栈顶暖，栈顶就等到了答案。
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{} // 还没等到更暖天气的下标
	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}
