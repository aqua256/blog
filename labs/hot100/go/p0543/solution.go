package p0543

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	// depth 返回以 node 为端点向下的最长链节点数；顺便用 左+右 更新经过 node 的直径
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := depth(node.Left), depth(node.Right)
		ans = max(ans, l+r)
		return max(l, r) + 1
	}
	depth(root)
	return ans
}
