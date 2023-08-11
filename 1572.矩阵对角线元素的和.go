/*
 * @lc app=leetcode.cn id=1572 lang=golang
 *
 * [1572] 矩阵对角线元素的和
 */

// @lc code=start
func diagonalSum(mat [][]int) (sum int) {
	n := len(mat)
	for i := 0; i < n; i++ {
		if i == n-i-1 {
			sum += mat[i][i]
		} else {
			sum += mat[i][i] + mat[i][n-i-1]
		}
	}
	return
}

// @lc code=end
