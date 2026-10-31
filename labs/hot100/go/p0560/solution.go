package p0560

func subarraySum(nums []int, k int) int {
	cnt := map[int]int{0: 1} // 前缀和 -> 出现次数；空前缀和为 0
	sum, ans := 0, 0
	for _, x := range nums {
		sum += x
		ans += cnt[sum-k] // 以当前位置结尾、和为 k 的子数组个数
		cnt[sum]++
	}
	return ans
}
