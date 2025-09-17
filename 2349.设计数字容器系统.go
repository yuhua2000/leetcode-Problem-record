package leetcode

import (
	"container/heap"
)

// IntHeap 为最小堆（基于 container/heap）
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type NumberContainers struct {
	index       map[int]int
	numberIndex map[int]*IntHeap
}

func Constructor() NumberContainers {
	return NumberContainers{
		index:       make(map[int]int),
		numberIndex: make(map[int]*IntHeap),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	this.index[index] = number

	if _, ok := this.numberIndex[number]; !ok {
		this.numberIndex[number] = &IntHeap{}
	}

	heap.Push(this.numberIndex[number], index)

}

func (this *NumberContainers) Find(number int) int {
	h, ok := this.numberIndex[number]
	if !ok {
		return -1
	}

	for h.Len() > 0 && this.index[(*h)[0]] != number {
		heap.Pop(h)
	}

	if h.Len() == 0 {
		return -1
	}

	return (*h)[0]
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
