/*
 * @lc app=leetcode.cn id=2413 lang=golang
 *
 * [2413] 最小偶倍数
 */

// @lc code=start
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}

// @lc code=end

