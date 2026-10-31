package p0105

func buildTree(preorder []int, inorder []int) *TreeNode {
	pos := make(map[int]int, len(inorder)) // 值 -> 中序下标，O(1) 切分
	for i, v := range inorder {
		pos[v] = i
	}
	// 前序首个是根；根在中序的位置把区间切成左右子树，左子树大小 = k-inL
	var build func(preL, inL, n int) *TreeNode
	build = func(preL, inL, n int) *TreeNode {
		if n == 0 {
			return nil
		}
		root := &TreeNode{Val: preorder[preL]}
		leftSize := pos[root.Val] - inL
		root.Left = build(preL+1, inL, leftSize)
		root.Right = build(preL+1+leftSize, inL+leftSize+1, n-1-leftSize)
		return root
	}
	return build(0, 0, len(preorder))
}
