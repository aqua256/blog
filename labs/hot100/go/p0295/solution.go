package p0295

import "container/heap"

// 两个堆把数据劈成两半：low 是较小一半（大顶堆），high 是较大一半（小顶堆）。
// 约定 len(low) == len(high) 或 len(low) == len(high)+1。
type MedianFinder struct {
	low  *intHeap // 存相反数，模拟大顶堆
	high *intHeap
}

func Constructor() MedianFinder {
	return MedianFinder{low: &intHeap{}, high: &intHeap{}}
}

func (mf *MedianFinder) AddNum(num int) {
	// 先进 low，再把 low 的最大值挪到 high，保证 low 全部 <= high
	heap.Push(mf.low, -num)
	heap.Push(mf.high, -heap.Pop(mf.low).(int))
	// high 多了就还一个给 low
	if mf.high.Len() > mf.low.Len() {
		heap.Push(mf.low, -heap.Pop(mf.high).(int))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.low.Len() > mf.high.Len() {
		return float64(-(*mf.low)[0])
	}
	return float64(-(*mf.low)[0]+(*mf.high)[0]) / 2
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
