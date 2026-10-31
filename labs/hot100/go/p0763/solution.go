package p0763

func partitionLabels(s string) []int {
	var last [26]int // 每个字母最后出现的位置
	for i := range len(s) {
		last[s[i]-'a'] = i
	}
	var res []int
	start, end := 0, 0
	for i := range len(s) {
		end = max(end, last[s[i]-'a']) // 当前段至少要延伸到这里
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}
