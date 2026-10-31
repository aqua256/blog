package p0042

func trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax, water := 0, 0, 0
	for l < r {
		leftMax = max(leftMax, height[l])
		rightMax = max(rightMax, height[r])
		// 哪边的已知最高更矮，哪边的水位就已确定：对面一定有更高的墙兜底
		if leftMax < rightMax {
			water += leftMax - height[l]
			l++
		} else {
			water += rightMax - height[r]
			r--
		}
	}
	return water
}
