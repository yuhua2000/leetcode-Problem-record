package leetcode

import "container/heap"

/*
 * @lc app=leetcode.cn id=1705 lang=golang
 *
 * [1705] 吃苹果的最大数目
 */

// @lc code=start

type Apple struct {
	Quantity int
	Expiry   int
}

type Apples []Apple

func (a Apples) Len() int {
	return len(a)
}

func (a Apples) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Apples) Less(i, j int) bool {
	return a[i].Expiry < a[j].Expiry
}

func (a *Apples) Push(x interface{}) {
	*a = append(*a, x.(Apple))
}

func (a *Apples) Pop() interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[0 : n-1]
	return x
}

func eatenApples(apples []int, days []int) int {
	var applesExpiry *Apples = new(Apples)
	result := 0
	for i := 0; i < len(apples) || applesExpiry.Len() > 0; i++ {
		if i < len(apples) && apples[i] > 0 {
			heap.Push(applesExpiry, Apple{
				Quantity: apples[i],
				Expiry:   i + days[i],
			})
		}

		for applesExpiry.Len() > 0 {
			appleExpiry := heap.Pop(applesExpiry).(Apple)
			if appleExpiry.Expiry > i {
				appleExpiry.Quantity--
			} else {
				continue
			}
			result++
			if appleExpiry.Quantity > 0 {
				heap.Push(applesExpiry, appleExpiry)
			}
			break
		}
	}

	return result
}

// @lc code=end
