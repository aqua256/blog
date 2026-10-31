package p0075

// 荷兰国旗：[0,p0) 全是 0，[p0,i) 全是 1，(p2,n-1] 全是 2，[i,p2] 待处理。
func sortColors(nums []int) {
	p0, i, p2 := 0, 0, len(nums)-1
	for i <= p2 {
		switch nums[i] {
		case 0:
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++ // 换过来的一定是 1（或 i==p0 时的自己），可以前进
		case 1:
			i++
		case 2:
			nums[i], nums[p2] = nums[p2], nums[i]
			p2-- // 换过来的还没看过，i 不动
		}
	}
}
