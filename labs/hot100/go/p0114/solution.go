package p0114

// O(1) 空间：把左子树整体接到右边，原右子树挂到左子树最右节点（前序前驱）之后
func flatten(root *TreeNode) {
	for cur := root; cur != nil; cur = cur.Right {
		if cur.Left == nil {
			continue
		}
		pre := cur.Left
		for pre.Right != nil {
			pre = pre.Right
		}
		pre.Right = cur.Right
		cur.Right, cur.Left = cur.Left, nil
	}
}
