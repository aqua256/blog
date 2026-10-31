package tpl_sliding_window

// 适用前提：窗口合法性对「扩张/收缩」单调——合法窗口缩小仍合法（求最长），
// 或合法窗口扩大仍合法（求最短）。负数等破坏单调性时改用前缀和。

// 1) 最长合法窗口：右端扩张，不合法就收左端，收到合法后更新答案。
func longestValid(s string, maxRepeat int) int {
	cnt := map[byte]int{}
	best, left := 0, 0
	for right := 0; right < len(s); right++ {
		cnt[s[right]]++
		for cnt[s[right]] > maxRepeat { // 只有新进来的字符可能破坏合法性
			cnt[s[left]]--
			left++
		}
		best = max(best, right-left+1)
	}
	return best
}

// 2) 最短合法窗口：右端扩张，一旦合法就尽量收左端，收的过程中更新答案。
// 例：和 >= target 的最短子数组（nums 全为正）。
func shortestValid(nums []int, target int) int {
	best, sum, left := len(nums)+1, 0, 0
	for right, x := range nums {
		sum += x
		for sum >= target {
			best = min(best, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if best > len(nums) {
		return 0
	}
	return best
}

// 3) 定长窗口：先填满 k 个，之后每步进一个、出一个。
func maxSumFixed(nums []int, k int) int {
	sum := 0
	for _, x := range nums[:k] {
		sum += x
	}
	best := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		best = max(best, sum)
	}
	return best
}
