/*
 * @lc app=leetcode.cn id=1253 lang=golang
 *
 * [1253] 重构 2 行二进制矩阵
 */

// @lc code=start
func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	sumVal := 0
	twoNum := 0
	for i := 0; i < n; i++ {
		if colsum[i] == 2 {
			twoNum++
		}
		sumVal += colsum[i]
	}
	if sumVal != upper+lower || twoNum > min(upper, lower) {
		return [][]int{}
	}
	upper -= twoNum
	lower -= twoNum
	res := make([][]int, 2)
	res[0] = make([]int, n)
	res[1] = make([]int, n)
	for i := 0; i < n; i++ {
		if colsum[i] == 2 {
			res[0][i] = 1
			res[1][i] = 1
		} else if colsum[i] == 1 {
			if upper > 0 {
				upper--
				res[0][i] = 1
			} else {
				res[1][i] = 1
			}
		}
	}
	return res
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
