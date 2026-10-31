package p0101

func isSymmetric(root *TreeNode) bool {
	// 两棵子树互为镜像：根相等，且外侧对外侧、内侧对内侧
	var mirror func(a, b *TreeNode) bool
	mirror = func(a, b *TreeNode) bool {
		if a == nil || b == nil {
			return a == b
		}
		return a.Val == b.Val && mirror(a.Left, b.Right) && mirror(a.Right, b.Left)
	}
	return mirror(root.Left, root.Right)
}
