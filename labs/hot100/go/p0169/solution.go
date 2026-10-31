package p0169

// Boyer-Moore 投票：不同的两个数互相抵消，众数个数 > n/2，抵消不完，最后留下的候选就是它。
func majorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, x := range nums {
		if count == 0 {
			candidate = x
		}
		if x == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
