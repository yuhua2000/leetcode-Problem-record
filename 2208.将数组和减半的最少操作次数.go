/*
 * @lc app=leetcode.cn id=2208 lang=golang
 *
 * [2208] 将数组和减半的最少操作次数
 */

// @lc code=start

type PriorityQueue []float64

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(float64))
}

func (pq *PriorityQueue) Pop() any {
	x := (*pq)[pq.Len()-1]
	*pq = (*pq)[0 : pq.Len()-1]
	return x
}

func halveArray(nums []int) (res int) {
	sum := 0
	pq := &PriorityQueue{}
	for _, n := range nums {
		heap.Push(pq, float64(n))
		sum += n
	}
	half := float64(sum) / 2
	for half > 0 {
		x := heap.Pop(pq).(float64) / 2
		half -= x
		res++
		heap.Push(pq, x)
	}
	return res
}

// @lc code=end
