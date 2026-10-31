package tpl_two_pointers

// 相向双指针：有序数组 / 两端各有单调性。每步根据比较结果排除一整行候选。
func twoSumSorted(nums []int, target int) (int, int) {
	l, r := 0, len(nums)-1
	for l < r {
		switch s := nums[l] + nums[r]; {
		case s < target:
			l++ // nums[l] 和任何 r 配都不够大，丢掉 l
		case s > target:
			r--
		default:
			return l, r
		}
	}
	return -1, -1
}

// 同向快慢指针：原地过滤。[0, slow) 是保留区，fast 负责扫描。
func filterInPlace(nums []int, keep func(int) bool) int {
	slow := 0
	for fast := range nums {
		if keep(nums[fast]) {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow // 新长度
}
