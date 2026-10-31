package p0215

import "container/heap"

func findKthLargest(nums []int, k int) int {
	// 维护大小为 k 的小顶堆：堆里是目前最大的 k 个，堆顶就是第 k 大
	h := &minHeap{}
	for _, x := range nums {
		if h.Len() < k {
			heap.Push(h, x)
		} else if x > (*h)[0] {
			(*h)[0] = x
			heap.Fix(h, 0)
		}
	}
	return (*h)[0]
}

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
