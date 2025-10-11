package leetcode

/*
 * @lc app=leetcode.cn id=2520 lang=golang
 *
 * [2520] 统计能整除数字的位数
 */

// @lc code=start
func countDigits(num int) int {
	result := 0
	for i := num; i > 0; i /= 10 {
		if num%(i%10) == 0 {
			result++
		}
	}
	return result
}

// @lc code=end
