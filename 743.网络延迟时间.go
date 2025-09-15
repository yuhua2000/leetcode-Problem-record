package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=743 lang=golang
 *
 * [743] 网络延迟时间
 */

// @lc code=start
func networkDelayTime(times [][]int, n int, k int) int {
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

// @lc code=end
