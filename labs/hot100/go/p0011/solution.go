package p0011

func maxArea(height []int) int {
	l, r, best := 0, len(height)-1, 0
	for l < r {
		best = max(best, min(height[l], height[r])*(r-l))
		// 短板决定上限：移动长板只会让宽变小、高不变大，所以只能移短板
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return best
}
