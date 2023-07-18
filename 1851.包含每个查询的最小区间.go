/*
 * @lc app=leetcode.cn id=1851 lang=golang
 *
 * [1851] 包含每个查询的最小区间
 */

// @lc code=start

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() any {
	n, old := len(*pq), *pq
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func (pq PriorityQueue) Top() []int {
	return pq[0]

}

func minInterval(intervals [][]int, queries []int) []int {
	qindex := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		qindex[i] = i
	}

	sort.Slice(qindex, func(i, j int) bool {
		return queries[qindex[i]] < queries[qindex[j]]
	})

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	pq := &PriorityQueue{}
	res := make([]int, len(queries))
	i := 0
	for _, qi := range qindex {
		for ; i < len(intervals) && intervals[i][0] <= queries[qi]; i++ {
			heap.Push(pq, []int{intervals[i][1] - intervals[i][0] + 1, intervals[i][0], intervals[i][1]})
		}

		for pq.Len() > 0 && pq.Top()[2] < queries[qi] {
			heap.Pop(pq)
		}

		if pq.Len() > 0 {
			res[qi] = pq.Top()[0]
		} else {
			res[qi] = -1
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
