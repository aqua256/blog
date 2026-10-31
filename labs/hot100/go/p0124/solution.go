package p0124

import "math"

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	// gain 返回从 node 向下的单链最大和；负贡献直接丢掉（取 0）
	var gain func(node *TreeNode) int
	gain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := max(gain(node.Left), 0)
		r := max(gain(node.Right), 0)
		ans = max(ans, node.Val+l+r) // 以 node 为拐点的路径
		return node.Val + max(l, r)  // 交给父节点只能选一边
	}
	gain(root)
	return ans
}
