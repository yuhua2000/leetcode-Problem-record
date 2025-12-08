package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=1925 lang=golang
 *
 * [1925] 统计平方和三元组的数目
 */

// @lc code=start
func countTriples(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			sumOfSquares := i*i + j*j
			c := int(math.Sqrt(float64(sumOfSquares + 1)))
			if c <= n && c*c == sumOfSquares {
				result++
			}
		}
	}
	return result
}

// @lc code=end
