package p0142

// ListNode 是 LeetCode 预置的单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// fromSlice 由切片构造链表（测试辅助）。
func fromSlice(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// toSlice 把链表转回切片（测试辅助，无环链表）。
func toSlice(head *ListNode) []int {
	out := []int{}
	for ; head != nil; head = head.Next {
		out = append(out, head.Val)
	}
	return out
}
