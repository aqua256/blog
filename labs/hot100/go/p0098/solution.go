package p0098

import "math"

func isValidBST(root *TreeNode) bool {
	// 每个节点必须落在祖先给出的开区间 (lo, hi) 内，而不只是和父节点比
	var valid func(node *TreeNode, lo, hi int) bool
	valid = func(node *TreeNode, lo, hi int) bool {
		if node == nil {
			return true
		}
		if node.Val <= lo || node.Val >= hi {
			return false
		}
		return valid(node.Left, lo, node.Val) && valid(node.Right, node.Val, hi)
	}
	return valid(root, math.MinInt, math.MaxInt)
}
