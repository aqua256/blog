package p0153

// 以末尾元素为界：> last 的在左段，<= last 的在右段（最小值是右段第一个）。
// 谓词「nums[i] <= last」是 false...true，二分找第一个 true。
func findMin(nums []int) int {
	last := nums[len(nums)-1]
	lo, hi := 0, len(nums)-1 // 末尾本身一定满足，不必纳入搜索
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] <= last {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return nums[lo]
}
