package leetcode

import (
	"container/heap"
)

type NextHeap []int

func (n NextHeap) Len() int           { return len(n) }
func (n NextHeap) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n NextHeap) Less(i, j int) bool { return n[i] < n[j] }

func (n *NextHeap) Push(v any) {
	*n = append(*n, v.(int))
}

func (n *NextHeap) Pop() any {
	x := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return x
}

func maximumJumps1(nums []int, target int) int {
	jumpNum := make([]int, len(nums))
	for i := range jumpNum {
		jumpNum[i] = -1
	}
	visit := map[int]bool{0: true}
	jumpNum[0] = 0

	h := &NextHeap{}
	heap.Push(h, 0)

	for h.Len() > 0 {
		nowIndex := heap.Pop(h).(int)
		now := nums[nowIndex]
		for nextIndex := nowIndex + 1; nextIndex < len(nums); nextIndex++ {
			next := nums[nextIndex]
			if next-now <= target && next-now >= -target {
				jumpNum[nextIndex] = max(jumpNum[nextIndex], jumpNum[nowIndex]+1)
				if !visit[nextIndex] {
					heap.Push(h, nextIndex)
					visit[nextIndex] = true

				}
			}
		}
	}

	return jumpNum[len(nums)-1]
}

func maximumJumps(nums []int, target int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if abs(nums[i]-nums[j]) <= target && dp[j] != -1 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(nums)-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
输入：nums = [1,3,6,4,1,2], target = 2
输出：3
解释：要想以最大跳跃次数从下标 0 到下标 n - 1 ，可以按下述跳跃序列执行操作：
- 从下标 0 跳跃到下标 1 。
- 从下标 1 跳跃到下标 3 。
- 从下标 3 跳跃到下标 5 。
可以证明，从 0 到 n - 1 的所有方案中，不存在比 3 步更长的跳跃序列。因此，答案是 3 。

一个全局map 记录到这个节点 最多需要多少次跳跃

从 0 出发 开始触发，可以跳跃到下表 进入堆

出堆

堆是空是结束

最后看map里面 最后一个下标的值
*/
