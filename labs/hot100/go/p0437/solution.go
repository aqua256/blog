package p0437

func pathSum(root *TreeNode, targetSum int) int {
	// 前缀和：根到当前节点的和 s，若祖先处出现过 s-target，中间那段就是一条路径
	cnt := map[int]int{0: 1}
	ans := 0
	var dfs func(node *TreeNode, s int)
	dfs = func(node *TreeNode, s int) {
		if node == nil {
			return
		}
		s += node.Val
		ans += cnt[s-targetSum]
		cnt[s]++
		dfs(node.Left, s)
		dfs(node.Right, s)
		cnt[s]-- // 回溯：离开子树后，这个前缀不再是祖先
	}
	dfs(root, 0)
	return ans
}
