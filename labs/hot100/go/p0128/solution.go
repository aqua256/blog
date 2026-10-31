package p0128

func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))
	for _, x := range nums {
		set[x] = true
	}
	best := 0
	for x := range set { // 遍历去重后的集合，避免重复起点退化成 O(n^2)
		if set[x-1] {
			continue // 只从序列起点开始数
		}
		y := x
		for set[y+1] {
			y++
		}
		best = max(best, y-x+1)
	}
	return best
}
