package p0076

func minWindow(s string, t string) string {
	var need [128]int
	for i := range len(t) {
		need[t[i]]++
	}
	missing := len(t) // 还差多少个字符（计重数）才能覆盖 t
	start, best := 0, len(s)+1
	left := 0
	for right := 0; right < len(s); right++ {
		if need[s[right]] > 0 {
			missing--
		}
		need[s[right]]-- // 负数表示窗口里多出来的
		for missing == 0 {
			if right-left+1 < best {
				start, best = left, right-left+1
			}
			need[s[left]]++
			if need[s[left]] > 0 {
				missing++ // 移走的是必需字符，窗口不再合法
			}
			left++
		}
	}
	if best > len(s) {
		return ""
	}
	return s[start : start+best]
}
