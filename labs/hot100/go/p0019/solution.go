package p0019

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	// fast 先走 n+1 步，二者间隔 n 个节点；fast 到 nil 时 slow 正好在待删节点前驱
	fast, slow := dummy, dummy
	for range n + 1 {
		fast = fast.Next
	}
	for fast != nil {
		fast, slow = fast.Next, slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
