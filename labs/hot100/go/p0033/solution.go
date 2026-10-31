package p0033

// 把问题改写成「找第一个位于 target 处或其右侧的下标」，谓词仍是 false...true。
// 以末尾 last 为界判断 nums[i] 和 target 各在哪一段，再比较。
func search(nums []int, target int) int {
	last := nums[len(nums)-1]
	isRight := func(i int) bool { // nums[i] 是否在 target 位置或其右侧
		x := nums[i]
		if x > last { // x 在左段
			return target > last && x >= target
		}
		// x 在右段：target 在左段则 x 一定在其右；否则同段直接比较
		return target > last || x >= target
	}
	lo, hi := 0, len(nums)-1 // 末尾一定满足 isRight
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isRight(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if nums[lo] != target {
		return -1
	}
	return lo
}
