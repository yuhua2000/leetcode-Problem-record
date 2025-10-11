package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1552 lang=golang
 *
 * [1552] 两球之间的磁力
 */

// @lc code=start
func maxDistance(position []int, m int) (result int) {
	sort.Ints(position)
	check := func(x int) bool {
		pre, cnt := position[0], 1
		for i := 1; i < len(position); i++ {
			if position[i]-pre >= x {
				pre = position[i]
				cnt++
			}
		}
		return cnt >= m
	}
	left, right := 1, position[len(position)-1]-position[0]
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return
}

// @lc code=end
