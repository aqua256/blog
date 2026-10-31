package p0034

func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	// 最后一个 <= target = 第一个 >= target+1 的前一个
	end := lowerBound(nums, target+1) - 1
	return []int{start, end}
}

// 左闭右开 [lo, hi)，返回第一个 >= target 的下标
func lowerBound(nums []int, target int) int {
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
