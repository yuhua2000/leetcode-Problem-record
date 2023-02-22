/*
 * @lc app=leetcode.cn id=1326 lang=golang
 *
 * [1326] 灌溉花园的最少水龙头数目
 */

// @lc code=start
// func minTaps(n int, ranges []int) int {
// 	intervals := make([][2]int, n+1)
// 	for i, r := range ranges {
// 		intervals[i] = [2]int{max(0, i-r), min(i+r, n)}
// 	}
// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] < intervals[j][0]
// 	})
// 	dp := make([]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		dp[i] = math.MaxInt
// 	}
// 	dp[0] = 0
// 	for _, p := range intervals {
// 		start, end := p[0], p[1]
// 		if dp[start] == math.MaxInt {
// 			return -1
// 		}
// 		for j := start; j <= end; j++ {
// 			dp[j] = min(dp[j], 1+dp[start])
// 		}
// 	}
// 	return dp[n]
// }
func minTaps(n int, ranges []int) int {
	rightMost := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		rightMost[i] = i
	}
	for i, r := range ranges {
		start := max(0, i-r)
		end := min(n, i+r)
		rightMost[start] = max(rightMost[start], end)
	}
	last, ret, pre := 0, 0, 0
	for i := 0; i < n; i++ {
		last = max(last, rightMost[i])
		if last == i {
			return -1
		}
		if i == pre {
			pre = last
			ret++
		}
	}
	return ret

}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
