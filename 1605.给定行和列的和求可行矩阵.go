/*
 * @lc app=leetcode.cn id=1605 lang=golang
 *
 * [1605] 给定行和列的和求可行矩阵
 */

// @lc code=start
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	i, j := 0, 0
	for i < n && j < m {
		v := min(rowSum[i], colSum[j])
		matrix[i][j] = v
		rowSum[i] -= v
		colSum[j] -= v
		if rowSum[i] == 0 {
			i++
		}
		if colSum[j] == 0 {
			j++
		}
	}
	return matrix
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

