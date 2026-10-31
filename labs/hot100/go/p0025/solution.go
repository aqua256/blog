package p0025

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy // 上一组的尾巴
	for {
		// 探路：不足 k 个就保持原样
		tail := prev
		for range k {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		next := tail.Next
		// 反转 [prev.Next, tail]，反转后原头变成组尾
		first := prev.Next
		pre, cur := next, first
		for cur != next {
			nxt := cur.Next
			cur.Next = pre
			pre, cur = cur, nxt
		}
		prev.Next = tail
		prev = first
	}
}
