package p0283

func moveZeroes(nums []int) {
	slow := 0 // [0, slow) 是已处理好的非零元素
	for fast := range nums {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
