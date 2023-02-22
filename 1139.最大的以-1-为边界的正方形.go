/*
 * @lc app=leetcode.cn id=1139 lang=golang
 *
 * [1139] 最大的以 1 为边界的正方形
 */

// @lc code=start
func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	left := make([][]int, m+1) // 扩容一个处理边界问题
	up := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		left[i] = make([]int, n+1)
		up[i] = make([]int, n+1)
	}
	maxBorder := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == 1 {
				left[i][j] = left[i][j-1] + 1
				up[i][j] = up[i-1][j] + 1
				border := min(left[i][j], up[i][j])
				for left[i-border+1][j] < border || up[i][j-border+1] < border {
					border--
				}
				maxBorder = max(maxBorder, border)
			}
		}
	}
	return maxBorder * maxBorder

}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
