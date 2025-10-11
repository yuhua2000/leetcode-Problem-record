package leetcode

/*
 * @lc app=leetcode.cn id=1186 lang=golang
 *
 * [1186] 删除一次得到子数组最大和
 */

// @lc code=start
func maximumSum(arr []int) int {
	dp0, dp1, res := arr[0], 0, arr[0]
	for i := 1; i < len(arr); i++ {
		dp0, dp1 = max(dp0, 0)+arr[i], max(dp1+arr[i], dp0)
		res = max(res, max(dp0, dp1))
	}
	return res
}

// @lc code=end
