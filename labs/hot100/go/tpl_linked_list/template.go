package tpl_linked_list

// 哑节点：头节点可能被删/被换时，统一成「改 prev.Next」
func removeVal(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for prev := dummy; prev.Next != nil; {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return dummy.Next
}

// 原地反转：prev / cur / next 三指针
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}

// 快慢指针找中点：奇数返回正中，偶数返回后半段第一个
func middle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
