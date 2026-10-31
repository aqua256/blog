package p0003

func lengthOfLongestSubstring(s string) int {
	last := map[byte]int{} // 字符 -> 最近一次出现的下标
	best, left := 0, 0
	for right := 0; right < len(s); right++ {
		if i, ok := last[s[right]]; ok && i >= left {
			left = i + 1 // 直接跳过重复字符，左端只增不减
		}
		last[s[right]] = right
		best = max(best, right-left+1)
	}
	return best
}
