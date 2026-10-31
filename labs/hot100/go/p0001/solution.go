package p0001

func twoSum(nums []int, target int) []int {
	seen := map[int]int{} // 值 -> 下标
	for i, x := range nums {
		if j, ok := seen[target-x]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return nil
}
