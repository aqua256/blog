package p0160

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 两指针各走 a+c+b 步后必然同时到达交点（或同时为 nil）
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
