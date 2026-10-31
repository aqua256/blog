package tpl_binary_tree

// 1) 后序 DFS：子树把「向上可用的信息」返回给父节点，
//    「跨越当前节点」的答案在外层变量里更新（直径 / 最大路径和 / LCA 同构）
func postorder(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0 // 空树的基值，按题意定
		}
		l, r := dfs(node.Left), dfs(node.Right)
		ans = max(ans, l+r)  // 以 node 为拐点的答案
		return max(l, r) + 1 // 父节点只能接一条链
	}
	dfs(root)
	return ans
}

// 2) 前序 DFS：从上往下携带状态（深度 / 区间上下界 / 前缀和）
func preorder(root *TreeNode) []int {
	res := []int{}
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) { // 每层第一次到达
			res = append(res, node.Val)
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return res
}

// 3) BFS 层序：队列一次处理完整一层
func bfs(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	for q := []*TreeNode{root}; len(q) > 0; {
		var next []*TreeNode
		level := []int{}
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
