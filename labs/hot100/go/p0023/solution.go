package p0023

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	h := &nodeHeap{}
	for _, l := range lists {
		if l != nil {
			*h = append(*h, l)
		}
	}
	heap.Init(h)
	dummy := &ListNode{}
	tail := dummy
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode) // 当前 k 个头里最小的
		tail.Next = node
		tail = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}

// nodeHeap 按 Val 排序的小顶堆
type nodeHeap []*ListNode

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
