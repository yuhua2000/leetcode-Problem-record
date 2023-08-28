/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		n := len(ans) - 1
		if intervals[i][1] <= ans[n][1] {
			continue
		} else if intervals[i][0] <= ans[n][1] {
			ans[n][1] = intervals[i][1]
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

// @lc code=end
