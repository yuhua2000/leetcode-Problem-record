/*
 * @lc app=leetcode.cn id=1014 lang=golang
 *
 * [1014] 最佳观光组合
 */
// @lc code=start
func maxScoreSightseeingPair(values []int) int {
	ans := 0
	mx := values[0] + 0
	for i := 1; i < len(values); i++ {
		ans = max(ans, mx+values[i]-i)
		mx = max(mx, values[i]+i)
	}
	return ans
}

// @lc code=end
