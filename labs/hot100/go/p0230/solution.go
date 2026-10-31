package p0230

// 迭代中序，第 k 个弹出的就是答案，可以提前停
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for cur := root; ; {
		for ; cur != nil; cur = cur.Left {
			stack = append(stack, cur)
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k--; k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
}
