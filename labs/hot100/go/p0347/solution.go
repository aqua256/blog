package p0347

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	for _, x := range nums {
		freq[x]++
	}
	// 按频次的小顶堆，超过 k 个就弹出最小的，剩下的就是前 k 高
	h := &pairHeap{}
	for num, cnt := range freq {
		heap.Push(h, pair{num, cnt})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, 0, k)
	for _, p := range *h {
		res = append(res, p.num)
	}
	return res
}

type pair struct{ num, cnt int }

type pairHeap []pair

func (h pairHeap) Len() int           { return len(h) }
func (h pairHeap) Less(i, j int) bool { return h[i].cnt < h[j].cnt }
func (h pairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(x any)        { *h = append(*h, x.(pair)) }
func (h *pairHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
