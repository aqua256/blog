package p0142

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			// 相遇后，一个从头走、一个从相遇点走，再次相遇即入环点
			for p := head.Next; p != slow; {
				p, slow = p.Next, slow.Next
			}
			return slow
		}
	}
	return nil
}
