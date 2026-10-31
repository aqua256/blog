package p0056

import "slices"

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
	ans := [][]int{intervals[0]}
	for _, in := range intervals[1:] {
		last := ans[len(ans)-1]
		if in[0] <= last[1] { // 有重叠（端点相接也算），延长右端
			last[1] = max(last[1], in[1])
		} else {
			ans = append(ans, in)
		}
	}
	return ans
}
