package p0015

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	var ans [][]int
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break // 最小的都 > 0，不可能凑出 0
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue // 第一个数去重
		}
		l, r := i+1, n-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				l++
			case s > 0:
				r--
			default:
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l++; l < r && nums[l] == nums[l-1]; l++ { // 第二个数去重
				}
				r--
			}
		}
	}
	return ans
}
