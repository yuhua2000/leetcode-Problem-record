/*
 * @lc app=leetcode.cn id=2485 lang=golang
 *
 * [2485] 找出中枢整数
 */

// @lc code=start
func pivotInteger(n int) int {
	t := (n*n + n) / 2
	x := int(math.Sqrt(float64(t)))
	if x*x == t {
		return x
	}
	return -1
}

// @lc code=end

