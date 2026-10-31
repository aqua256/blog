package p0024

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		a, b := prev.Next, prev.Next.Next
		// prev -> a -> b -> next  变成  prev -> b -> a -> next
		a.Next = b.Next
		b.Next = a
		prev.Next = b
		prev = a
	}
	return dummy.Next
}
