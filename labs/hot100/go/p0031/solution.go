package p0031

// 1. 从右往左找第一个「下降点」i：nums[i] < nums[i+1]，其右侧是非增序列。
// 2. 若存在，在右侧找最右一个 > nums[i] 的 j，交换——让第 i 位增大得最少。
// 3. 把 i 右侧反转成升序（最小的后缀）。若不存在 i，整体反转即回到最小排列。
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
