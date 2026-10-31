package tpl_binary_search

// 统一写法：左闭右开 [lo, hi)。
// 不变量：[0, lo) 全部 < target，[hi, n) 全部 >= target。
// 循环结束时 lo == hi，即第一个 >= target 的下标（不存在则为 n）。
func lowerBound(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			hi = mid // mid 满足条件，答案在 [lo, mid]
		} else {
			lo = mid + 1 // mid 不满足，答案在 [mid+1, hi)
		}
	}
	return lo
}

// 整数数组上，其余三种边界都能由 lowerBound 推出：
//
//	第一个 >  x：lowerBound(x+1)
//	最后一个 < x：lowerBound(x) - 1
//	最后一个 <= x：lowerBound(x+1) - 1
func upperBound(nums []int, target int) int { return lowerBound(nums, target+1) }

func lastLess(nums []int, target int) int { return lowerBound(nums, target) - 1 }

func lastLessEq(nums []int, target int) int { return lowerBound(nums, target+1) - 1 }

// 更一般地：只要谓词 ok(i) 在 [lo, hi) 上是 false...false true...true，
// 就能用同一个循环找第一个 true。旋转数组、找峰值、二分答案都是这一型。
func firstTrue(lo, hi int, ok func(int) bool) int {
	for lo < hi {
		mid := lo + (hi-lo)/2
		if ok(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
