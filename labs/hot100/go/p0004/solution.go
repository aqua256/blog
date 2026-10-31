package p0004

import "math"

// 在较短数组上二分「左半部分从 nums1 取几个」：左半共 k=(m+n+1)/2 个，
// nums1 取 i 个、nums2 取 j=k-i 个。谓词 nums1[i] >= nums2[j-1] 随 i 单调 false...true，
// 第一个 true 的 i 恰好满足交叉条件 nums1[i-1] <= nums2[j] 且 nums2[j-1] <= nums1[i]。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	k := (m + n + 1) / 2
	lo, hi := 0, m // 在 [0, m) 里找，找不到就是 m
	for lo < hi {
		i := lo + (hi-lo)/2
		if nums1[i] >= nums2[k-i-1] {
			hi = i
		} else {
			lo = i + 1
		}
	}
	i, j := lo, k-lo

	at := func(a []int, idx, fallback int) int {
		if idx < 0 || idx >= len(a) {
			return fallback
		}
		return a[idx]
	}
	leftMax := max(at(nums1, i-1, math.MinInt), at(nums2, j-1, math.MinInt))
	if (m+n)%2 == 1 {
		return float64(leftMax)
	}
	rightMin := min(at(nums1, i, math.MaxInt), at(nums2, j, math.MaxInt))
	return float64(leftMax+rightMin) / 2
}
