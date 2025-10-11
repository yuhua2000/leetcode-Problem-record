package leetcode

import "container/heap"

/*
 * @lc app=leetcode.cn id=1499 lang=golang
 *
 * [1499] 满足不等式的最大值
 */

// @lc code=start

type PriorityQueue [][]int

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(x, y int) bool {
	if p[x][0] != p[y][0] {
		return p[x][0] < p[y][0]
	}
	return p[x][1] < p[y][1]
}

func (p PriorityQueue) Swap(x, y int) {
	p[x], p[y] = p[y], p[x]
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.([]int))
}

func (p *PriorityQueue) Pop() any {
	n := len(*p)
	x := (*p)[n-1]
	*p = (*p)[:n-1]
	return x
}

func (pq PriorityQueue) Top() []int {
	return pq[0]
}

func findMaxValueOfEquation(points [][]int, k int) int {
	res := -0x3f3f3f3f
	pq := &PriorityQueue{}
	for _, p := range points {
		x, y := p[0], p[1]
		for pq.Len() > 0 && x-pq.Top()[1] > k {
			heap.Pop(pq)
		}
		if pq.Len() > 0 {
			res = max(res, x+y-pq.Top()[0])
		}
		heap.Push(pq, []int{x - y, x})
	}
	return res
}

// @lc code=end
