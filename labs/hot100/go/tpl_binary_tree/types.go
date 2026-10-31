package tpl_binary_tree

import "math/rand/v2"

// TreeNode 为 LeetCode 预置类型。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// null 在层序数组中表示空节点。
const null = -1 << 31

// build 由 LeetCode 风格层序数组（含 null）构造树。
func build(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	q := []*TreeNode{root}
	for i := 1; i < len(vals); {
		node := q[0]
		q = q[1:]
		for _, child := range []**TreeNode{&node.Left, &node.Right} {
			if i < len(vals) && vals[i] != null {
				*child = &TreeNode{Val: vals[i]}
				q = append(q, *child)
			}
			i++
		}
	}
	return root
}

// serialize 把树转回层序数组（去掉末尾 null），与 build 互逆。
func serialize(root *TreeNode) []int {
	res := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			res = append(res, null)
			continue
		}
		res = append(res, node.Val)
		q = append(q, node.Left, node.Right)
	}
	for len(res) > 0 && res[len(res)-1] == null {
		res = res[:len(res)-1]
	}
	return res
}

// randTree 生成 n 个节点、随机形状、值域 [lo,hi] 的树。
func randTree(r *rand.Rand, n, lo, hi int) *TreeNode {
	if n == 0 {
		return nil
	}
	k := r.IntN(n)
	return &TreeNode{
		Val:   lo + r.IntN(hi-lo+1),
		Left:  randTree(r, k, lo, hi),
		Right: randTree(r, n-1-k, lo, hi),
	}
}

// randDistinctTree 生成 n 个节点、值为 1..n 打乱后的树（值互不相同）。
func randDistinctTree(r *rand.Rand, n int) *TreeNode {
	vals := r.Perm(n)
	idx := 0
	var gen func(n int) *TreeNode
	gen = func(n int) *TreeNode {
		if n == 0 {
			return nil
		}
		k := r.IntN(n)
		node := &TreeNode{Val: vals[idx] + 1}
		idx++
		node.Left = gen(k)
		node.Right = gen(n - 1 - k)
		return node
	}
	return gen(n)
}

// randBST 把 n 个不同的随机值依次插入得到 BST。
func randBST(r *rand.Rand, n int) *TreeNode {
	var root *TreeNode
	for _, v := range r.Perm(3 * n)[:n] {
		p := &root
		for *p != nil {
			if v < (*p).Val {
				p = &(*p).Left
			} else {
				p = &(*p).Right
			}
		}
		*p = &TreeNode{Val: v}
	}
	return root
}

// nodes 返回前序下的全部节点。
func nodes(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	res := []*TreeNode{root}
	res = append(res, nodes(root.Left)...)
	return append(res, nodes(root.Right)...)
}
