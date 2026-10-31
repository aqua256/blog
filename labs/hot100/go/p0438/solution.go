package p0438

func findAnagrams(s string, p string) []int {
	k := len(p)
	if len(s) < k {
		return nil
	}
	var need, win [26]int
	for i := range k {
		need[p[i]-'a']++
		win[s[i]-'a']++
	}
	var ans []int
	if win == need { // Go 数组可直接 == 比较，O(26)
		ans = append(ans, 0)
	}
	for i := k; i < len(s); i++ { // 定长窗口：进一个、出一个
		win[s[i]-'a']++
		win[s[i-k]-'a']--
		if win == need {
			ans = append(ans, i-k+1)
		}
	}
	return ans
}
