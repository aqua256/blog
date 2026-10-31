package p0146

import (
	"math/rand/v2"
	"slices"
	"testing"
)

// naive 朴素 LRU：切片按使用时间排列，末尾最新。
type naive struct {
	cap  int
	keys []int
	vals map[int]int
}

func (n *naive) touch(k int) {
	i := slices.Index(n.keys, k)
	n.keys = append(slices.Delete(n.keys, i, i+1), k)
}

func (n *naive) get(k int) int {
	v, ok := n.vals[k]
	if !ok {
		return -1
	}
	n.touch(k)
	return v
}

func (n *naive) put(k, v int) {
	if _, ok := n.vals[k]; ok {
		n.vals[k] = v
		n.touch(k)
		return
	}
	n.vals[k] = v
	n.keys = append(n.keys, k)
	if len(n.keys) > n.cap {
		delete(n.vals, n.keys[0])
		n.keys = n.keys[1:]
	}
}

func TestExample(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	want := []int{1}
	got := []int{c.Get(1)}
	c.Put(3, 3)
	got = append(got, c.Get(2))
	c.Put(4, 4)
	got = append(got, c.Get(1), c.Get(3), c.Get(4))
	want = append(want, -1, -1, 3, 4)
	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 146))
	for range 300 {
		capacity := 1 + r.IntN(5)
		c := Constructor(capacity)
		ref := &naive{cap: capacity, vals: map[int]int{}}
		for op := range 200 {
			k := r.IntN(8)
			if r.IntN(2) == 0 {
				v := r.IntN(100)
				c.Put(k, v)
				ref.put(k, v)
			} else if got, want := c.Get(k), ref.get(k); got != want {
				t.Fatalf("cap=%d op#%d Get(%d) = %d, want %d", capacity, op, k, got, want)
			}
		}
	}
}
