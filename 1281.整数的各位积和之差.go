package leetcode

/*
 * @lc app=leetcode.cn id=1281 lang=golang
 *
 * [1281] 整数的各位积和之差
 */

// @lc code=start
func subtractProductAndSum(n int) int {
	m, s := 1, 0
	for n > 0 {
		x := n % 10
		n = n / 10
		m *= x
		s += x
	}
	return m - s
}

// @lc code=end
