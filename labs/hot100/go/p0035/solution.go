package p0035

// 插入位置 = 第一个 >= target 的下标，就是 lowerBound 本身。
func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
