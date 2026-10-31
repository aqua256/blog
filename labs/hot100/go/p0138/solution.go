package p0138

func copyRandomList(head *Node) *Node {
	// 第一遍：建 旧节点 -> 新节点 映射；第二遍：照映射连 Next / Random
	clone := map[*Node]*Node{nil: nil}
	for p := head; p != nil; p = p.Next {
		clone[p] = &Node{Val: p.Val}
	}
	for p := head; p != nil; p = p.Next {
		clone[p].Next = clone[p.Next]
		clone[p].Random = clone[p.Random]
	}
	return clone[head]
}
