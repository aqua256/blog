package tpl_heap

import "container/heap"

// container/heap 五件套：Len / Less / Swap 用值接收者，Push / Pop 用指针接收者
// Less 用 < 是小顶堆，改成 > 就是大顶堆
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// Top K 大：容量为 k 的小顶堆，堆顶是门槛，比门槛大才进
func topK(nums []int, k int) []int {
	h := &minHeap{}
	for _, x := range nums {
		heap.Push(h, x) // 必须用 heap.Push，不是 h.Push
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return *h
}
