/*
 * @lc app=leetcode.cn id=1238 lang=golang
 *
 * [1238] 循环码排列
 */

// @lc code=start
func circularPermutation(n int, start int) []int {
	ans := make([]int, 1<<n)
	for i := 0; i < len(ans); i++ {
		ans[i] = (i >> 1) ^ i ^ start
	}
	return ans
}

// @lc code=end

