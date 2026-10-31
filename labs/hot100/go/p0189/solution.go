package p0189

import "slices"

func rotate(nums []int, k int) {
	k %= len(nums) // k 可能大于长度
	// 三次翻转：整体翻转后，前 k 个和后 n-k 个各自顺序反了，再分别翻回来
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
