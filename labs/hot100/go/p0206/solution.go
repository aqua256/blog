package p0206

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for cur := head; cur != nil; {
		next := cur.Next // 先存后继，再掉头
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}
