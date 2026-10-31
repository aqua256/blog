package p0148

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针切成两半：fast 从 head.Next 出发，保证两个节点时也能切开
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return merge(sortList(head), sortList(mid))
}

func merge(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next, a = a, a.Next
		} else {
			tail.Next, b = b, b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}
