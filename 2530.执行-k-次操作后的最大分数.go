package leetcode

import "container/heap"

/*
 * @lc app=leetcode.cn id=2530 lang=golang
 *
 * [2530] 执行 K 次操作后的最大分数
 */

// @lc code=start
func maxKelements(nums []int, k int) (result int64) {
	h := make(hp, 0)
	for _, num := range nums {
		heap.Push(&h, num)
	}
	for i := 0; i < k; i++ {
		num := heap.Pop(&h).(int)
		result += int64(num)
		heap.Push(&h, (num+2)/3)
	}
	return result
}

type hp []int

func (a hp) Len() int           { return len(a) }
func (a hp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a hp) Less(i, j int) bool { return a[i] > a[j] }

func (a *hp) Push(x any) {
	*a = append(*a, x.(int))
}

func (a *hp) Pop() any {
	result := (*a)[a.Len()-1]
	*a = (*a)[:a.Len()-1]
	return result
}

// @lc code=end
