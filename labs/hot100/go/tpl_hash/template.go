package tpl_hash

// 一遍扫描 + 哈希表：边遍历边查「之前见过的、能和当前元素配对的东西」。
// key 的设计是核心：值本身（两数之和）、规范化形态（异位词排序串）、前缀和（和为 K 的子数组）。
func pairIndex(nums []int, target int) (int, int, bool) {
	seen := map[int]int{} // 值 -> 下标
	for i, x := range nums {
		if j, ok := seen[target-x]; ok { // 先查后存，保证不和自己配对
			return j, i, true
		}
		seen[x] = i
	}
	return -1, -1, false
}

// 分组：把元素映射到规范化 key，同 key 归一组。
func groupBy[T any, K comparable](items []T, key func(T) K) map[K][]T {
	groups := map[K][]T{}
	for _, it := range items {
		k := key(it)
		groups[k] = append(groups[k], it)
	}
	return groups
}
