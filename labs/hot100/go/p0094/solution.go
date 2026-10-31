package p0094

// 迭代中序：一路压左链，弹出即访问，再转向右子树
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	for cur := root; cur != nil || len(stack) > 0; {
		for ; cur != nil; cur = cur.Left {
			stack = append(stack, cur)
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}
