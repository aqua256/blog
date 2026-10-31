package p0234

func isPalindrome(head *ListNode) bool {
	// 快慢指针找中点：奇数长度时 slow 停在正中，偶数时停在后半段起点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// 反转后半段
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev, slow = slow, next
	}
	for p, q := head, prev; q != nil; p, q = p.Next, q.Next {
		if p.Val != q.Val {
			return false
		}
	}
	return true
}
