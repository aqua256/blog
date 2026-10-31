package p0138

// Node 是 LeetCode 预置的带随机指针的链表节点。
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// build 由 [[val, randomIndex]] 构造链表，randomIndex 为 -1 表示 nil（测试辅助）。
func build(spec [][2]int) (*Node, []*Node) {
	nodes := make([]*Node, len(spec))
	for i, s := range spec {
		nodes[i] = &Node{Val: s[0]}
	}
	for i, s := range spec {
		if i+1 < len(nodes) {
			nodes[i].Next = nodes[i+1]
		}
		if s[1] >= 0 {
			nodes[i].Random = nodes[s[1]]
		}
	}
	if len(nodes) == 0 {
		return nil, nil
	}
	return nodes[0], nodes
}

// collect 沿 Next 收集所有节点（测试辅助）。
func collect(head *Node) []*Node {
	var out []*Node
	for ; head != nil; head = head.Next {
		out = append(out, head)
	}
	return out
}
