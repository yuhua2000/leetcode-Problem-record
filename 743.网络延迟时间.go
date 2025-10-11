package leetcode

import (
	"container/heap"
	"math"
)

/*
 * @lc app=leetcode.cn id=743 lang=golang
 *
 * [743] 网络延迟时间
 */

// @lc code=start

func networkDelayTime1(times [][]int, n int, k int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[k] = 0

	for range n - 1 {
		for _, time := range times {
			if dp[time[0]] != math.MaxInt && dp[time[0]]+time[2] < dp[time[1]] {
				dp[time[1]] = dp[time[0]] + time[2]
			}
		}
	}

	dp = dp[1:]
	maxTiem := dp[0]
	for _, time := range dp {
		if time == math.MaxInt {
			return -1
		}
		maxTiem = max(maxTiem, time)
	}

	return maxTiem
}

type Item struct {
	node int
	dist int
}

type MinHeap []Item

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].dist < m[j].dist }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(item any) {
	*m = append(*m, item.(Item))
}

func (m *MinHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][][2]int, n+1)
	for _, time := range times {
		graph[time[0]] = append(graph[time[0]], [2]int{time[1], time[2]})
	}

	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	dist[k] = 0

	h := &MinHeap{}
	heap.Push(h, Item{node: k, dist: 0})

	for h.Len() > 0 {
		itme := heap.Pop(h).(Item)
		for _, neighbor := range graph[itme.node] {
			v, w := neighbor[0], neighbor[1]
			if dist[itme.node]+w < dist[v] {
				dist[v] = dist[itme.node] + w
				heap.Push(h, Item{node: v, dist: dist[v]})
			}
		}
	}

	maxTime := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		maxTime = max(maxTime, dist[i])
	}

	return maxTime
}

// @lc code=end
