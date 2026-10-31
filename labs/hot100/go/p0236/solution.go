package p0236

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 返回子树中找到的 p 或 q（或已确定的 LCA）；左右各找到一个时当前节点就是 LCA
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
