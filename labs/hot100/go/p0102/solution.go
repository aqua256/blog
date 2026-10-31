package p0102

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		// 当前队列恰好是一整层
		level := make([]int, 0, len(q))
		var next []*TreeNode
		for _, node := range q {
			level = append(level, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, level)
		q = next
	}
	return res
}
