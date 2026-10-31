package p0049

import "slices"

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, s := range strs {
		b := []byte(s)
		slices.Sort(b) // 排序后的串作为同组的统一 key
		groups[string(b)] = append(groups[string(b)], s)
	}
	ans := make([][]string, 0, len(groups))
	for _, g := range groups {
		ans = append(ans, g)
	}
	return ans
}
