package tpl_backtracking

import "slices"

// 组合/子集型：start 控制「只往后选」，避免同一组合不同顺序重复出现
func combine(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(start int)
	dfs = func(start int) {
		res = append(res, slices.Clone(path)) // 收集：按题意加条件（长度 / 和）；必须拷贝
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i]) // 做选择
			dfs(i + 1)                   // 可重复选同一元素则传 i
			path = path[:len(path)-1]    // 撤销选择
		}
	}
	dfs(0)
	return res
}

// 排列型：每层从 0 扫，用 used 标记已在路径上的元素
func permute(nums []int) [][]int {
	res := [][]int{}
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			res = append(res, slices.Clone(path))
			return
		}
		for i, v := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, v)
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return res
}
